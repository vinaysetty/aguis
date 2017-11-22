package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"mime"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/autograde/aguis/database"
	"github.com/autograde/aguis/logger"
	"github.com/autograde/aguis/web"
	"github.com/autograde/aguis/web/auth"
	"github.com/autograde/aguis/web/graphql"
	"github.com/graphql-go/graphql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"github.com/sirupsen/logrus"
)

func init() {
	mustAddExtensionType := func(ext, typ string) {
		if err := mime.AddExtensionType(ext, typ); err != nil {
			panic(err)
		}
	}

	// On Windows, mime types are read from the registry, which often has
	// outdated content types. This enforces that the correct mime types
	// are used on all platforms.
	mustAddExtensionType(".html", "text/html")
	mustAddExtensionType(".css", "text/css")
	mustAddExtensionType(".js", "application/javascript")
	mustAddExtensionType(".jsx", "application/javascript")
	mustAddExtensionType(".map", "application/json")
	mustAddExtensionType(".ts", "application/x-typescript")
}

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
		public   = flag.String("http.public", "public", "directory to server static files from")

		dbFile = flag.String("database.file", tempFile("ag.db"), "database file")

		baseURL = flag.String("service.url", "localhost", "service base url")

		fake = flag.Bool("provider.fake", false, "enable fake provider")
	)
	flag.Parse()

	l := logrus.New()
	l.Formatter = logger.NewDevFormatter(l.Formatter)

	entryPoint := filepath.Join(*public, "index.html")
	if !fileExists(entryPoint) {
		l.WithField("path", entryPoint).Warn("could not find file")
	}

	db, err := database.NewGormDB("sqlite3", *dbFile, database.Logger{Logger: l})
	if err != nil {
		l.WithError(err).Fatal("could not connect to db")
	}
	defer func() {
		if dbErr := db.Close(); dbErr != nil {
			l.WithError(dbErr).Warn("error closing database")
		}
	}()

	e := newServer(l) // Create server with echo
	//registerAPI(l, e, db, &bh)                     // Create REST API
	run(l, e, *httpAddr) // Start the server
}

func newServer(l *logrus.Logger) *echo.Echo {
	e := echo.New()
	e.Logger = web.EchoLogger{Logger: l}
	e.HideBanner = true
	e.Use(
		middleware.Recover(),
		web.Logger(l),
		middleware.Secure(),
	)

	return e
}

func enableProviders(l logrus.FieldLogger, baseURL string, fake bool) map[string]bool {
	enabled := make(map[string]bool)

	if ok := auth.EnableProvider(&auth.Provider{
		Name:          "github",
		KeyEnv:        "GITHUB_KEY",
		SecretEnv:     "GITHUB_SECRET",
		CallbackURL:   auth.GetCallbackURL(baseURL, "github"),
		StudentScopes: []string{"user", "repo", "delete_repo"}, // For testing, consider to push to master
		TeacherScopes: []string{"user", "repo", "delete_repo"},
	}, func(key, secret, callback string, scopes ...string) goth.Provider {
		return github.New(key, secret, callback, scopes...)
	}); ok {
		enabled["github"] = true
	} else {
		l.WithFields(logrus.Fields{
			"provider": "github",
			"enabled":  false,
		}).Warn("environment variables not set")
	}
	return enabled
}
func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func handler(w http.ResponseWriter, r *http.Request) {
	result := executeQuery(r.URL.Query().Get("query"), graphql.Schema)
	json.NewEncoder(w).Encode(result)
}

func registerGraphqlAPI(e *echo.Echo, db database.Database) {
	api := e.Group("/api/v1")

	//Graphql endpoints
	api.GET("/graphql", echo.WrapHandler(http.HandlerFunc(handler)))
}

func run(l logrus.FieldLogger, e *echo.Echo, httpAddr string) {
	go func() {
		srvErr := e.Start(httpAddr)
		if srvErr == http.ErrServerClosed {
			l.Warn("shutting down the server")
			return
		}
		l.WithError(srvErr).Fatal("could not start server")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		l.WithError(err).Fatal("failure during server shutdown")
	}
}

func tempFile(name string) string {
	return filepath.Join(os.TempDir(), name)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
