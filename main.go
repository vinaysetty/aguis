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
		grpcAddr        = flag.String("grpc.addr", ":8090", "gRPC listen address")
		dbFile          = flag.String("database.file", tempFile("ag.db"), "database file")
		enableTLS       = flag.Bool("enable_tls", false, "Use TLS - required for HTTP2.")
		tlsCertFilePath = flag.String("tls_cert_file", "misc/localhost.crt", "Path to the CRT/PEM file.")
		tlsKeyFilePath  = flag.String("tls_key_file", "misc/localhost.key", "Path to the private key file.")

		// flag variables for HTTP web server
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
		public   = flag.String("http.public", "public", "directory to server static files from")
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
		Addr:    *grpcAddr,
		Handler: http.HandlerFunc(handler),
	}

	// Start HTTP Web server
	go func() {
		web.NewWebServer(l, *public, *httpAddr)
	}()

	l.Infof("Starting server at %v with TLS: %t", *grpcAddr, *enableTLS)
	if *enableTLS {
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
