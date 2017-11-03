package main

import (
	"flag"
	"github.com/autograde/aguis/database"
	"github.com/autograde/aguis/grpcutil"
	"github.com/autograde/aguis/logger"
	pb "github.com/autograde/aguis/proto/_proto/aguis/library"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	var (
		httpAddr = flag.String("http.addr", ":8090", "HTTP listen address")
		//public   = flag.String("http.public", "public", "directory to server static files from")

		dbFile = flag.String("database.file", tempFile("ag.db"), "database file")

		//baseURL = flag.String("service.url", "localhost", "service base url")

		//fake = flag.Bool("provider.fake", false, "enable fake provider")
	)

	flag.Parse()

	l := logrus.New()
	l.Formatter = logger.NewDevFormatter(l.Formatter)

	db, err := database.NewGormDB("sqlite3", *dbFile, database.Logger{Logger: l})
	if err != nil {
		l.WithError(err).Fatal("could not connect to db")
	}

	defer func() {
		if dbErr := db.Close(); dbErr != nil {
			l.WithError(dbErr).Warn("error closing database")
		}
	}()

	grpcServer := grpc.NewServer()
	pb.RegisterAutograderServiceServer(grpcServer, grpcutil.NewAutograderService(db))

	wrappedServer := grpcweb.WrapServer(grpcServer)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHttp(resp, req)
	}

	httpServer := http.Server{
		Addr:    *httpAddr,
		Handler: http.HandlerFunc(handler),
	}

	l.Infof("Starting server. http port: %d", *httpAddr)
	if err := httpServer.ListenAndServe(); err != nil {
		l.WithError(err).Fatal("failed starting http server")
	}
}

func tempFile(name string) string {
	return filepath.Join(os.TempDir(), name)
}
