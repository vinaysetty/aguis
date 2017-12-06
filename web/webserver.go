package web

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func NewWebServer(l logrus.FieldLogger, public, httpAddr string) {

	entryPoint := filepath.Join(public, "index.html")
	if !fileExists(entryPoint) {
		l.WithField("path", entryPoint).Warn("could not find file")
	}

	e := newServer()
	registerFrontend(e, entryPoint, public)
	runWebServer(l, e, httpAddr)
}

func newServer() *echo.Echo {
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
