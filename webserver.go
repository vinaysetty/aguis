package main

import (
	"context"
	"flag"
	"github.com/autograde/aguis/logger"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
		public   = flag.String("http.public", "public", "directory to server static files from")

		//baseURL  = flag.String("service.url", "localhost", "service base url")

		//fake = flag.Bool("provider.fake", false, "enable fake provider")
	)

	flag.Parse()

	l := logrus.New()
	l.Formatter = logger.NewDevFormatter(l.Formatter)

	entryPoint := filepath.Join(*public, "index.html")
	if !fileExists(entryPoint) {
		l.WithField("path", entryPoint).Warn("could not find file")
	}

	e := newWebServer()
	registerFrontend(e, entryPoint, *public)
	runWebServer(l, e, *httpAddr)
}

func newWebServer() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	return e
}

func registerFrontend(e *echo.Echo, entryPoint, public string) {
	index := func(c echo.Context) error {
		return c.File(entryPoint)
	}
	e.GET("/app", index)
	e.GET("/app/*", index)

	// TODO: Whitelisted files only.
	e.Static("/", public)
}

func runWebServer(l logrus.FieldLogger, e *echo.Echo, httpAddr string) {
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

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
