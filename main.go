package main

import (
	"flag"
	"net/http"
	"os"
	"path/filepath"

	pb "github.com/autograde/aguis/ag"
	"github.com/autograde/aguis/database"
	"github.com/autograde/aguis/logger"
	"github.com/autograde/aguis/web"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8090", "HTTP listen address")
		//public          = flag.String("http.public", "public", "directory to server static files from")
		dbFile          = flag.String("database.file", tempFile("ag.db"), "database file")
		enableTls       = flag.Bool("enable_tls", true, "Use TLS - required for HTTP2.")
		tlsCertFilePath = flag.String("tls_cert_file", "misc/localhost.crt", "Path to the CRT/PEM file.")
		tlsKeyFilePath  = flag.String("tls_key_file", "misc/localhost.key", "Path to the private key file.")

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
	pb.RegisterAutograderServiceServer(grpcServer, web.NewAutograderService(db))
	wrappedServer := grpcweb.WrapServer(grpcServer)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHttp(resp, req)
	}

	httpServer := http.Server{
		Addr:    *httpAddr,
		Handler: http.HandlerFunc(handler),
	}

	l.Infof("Starting server at %v with TLS: %t", *httpAddr, *enableTls)
	if *enableTls {
		if err := httpServer.ListenAndServeTLS(*tlsCertFilePath, *tlsKeyFilePath); err != nil {
			l.WithError(err).Fatal("failed starting http2 server")
		}
	} else {
		if err := httpServer.ListenAndServe(); err != nil {
			l.WithError(err).Fatal("failed starting http server")
		}
	}
}

func tempFile(name string) string {
	return filepath.Join(os.TempDir(), name)
}
