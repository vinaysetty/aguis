package main

import (
	"flag"
	"github.com/autograde/aguis/database"
	"github.com/autograde/aguis/logger"
	library "github.com/autograde/aguis/proto/_proto/aguis/library"
	"github.com/autograde/aguis/web"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
	"os"
	"path/filepath"
)

type autograderService struct {
	db *database.GormDB
}

func (s *autograderService) GetUser(ctx context.Context, userQuery *library.GetUserRequest) (*library.User, error) {
	SetGrpcHeaderAndTrailer(ctx)
	user, err := web.GetUser(userQuery, s.db)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *autograderService) GetUsers(ctx context.Context, query *library.Void) (*library.UsersResponse, error) {
	users, err := web.GetUsers(s.db)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *autograderService) UpdateUser(ctx context.Context, u *library.UpdateUserRequest) (*library.User, error) {
	SetGrpcHeaderAndTrailer(ctx)
	user, err := web.UpdateUser(u, s.db)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *autograderService) GetCourses(ctx context.Context, q *library.Void) (*library.Courses, error) {
	SetGrpcHeaderAndTrailer(ctx)
	courses, err := web.ListCourses(s.db)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func main() {

	var (
		httpAddr = flag.String("http.addr", ":8090", "HTTP listen address")
		public   = flag.String("http.public", "public", "directory to server static files from")

		dbFile = flag.String("database.file", tempFile("ag.db"), "database file")

		//baseURL = flag.String("service.url", "localhost", "service base url")

		//fake = flag.Bool("provider.fake", false, "enable fake provider")
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

	grpcServer := grpc.NewServer()
	library.RegisterAutograderServiceServer(grpcServer, &autograderService{db: db})

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

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func SetGrpcHeaderAndTrailer(ctx context.Context) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))
	grpc.SetTrailer(ctx, metadata.Pairs("Post-Response-Metadata", "Is-sent-as-trailers-unary"))

}
