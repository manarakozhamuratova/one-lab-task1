package httpserver

import (
	"context"
	"net/http"
	"one-lab/config"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) NewServer(cfg config.Config, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           cfg.ServerAddress,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	logrus.Info("Server is started at port http://localhost:9090") // у тебя порт динамический, а лог статичный
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	logrus.Info("Server shutting down...")
	return s.httpServer.Shutdown(ctx)
}
