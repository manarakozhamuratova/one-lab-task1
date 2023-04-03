package httpserver

import (
	"net/http"
	"one-lab/config"
	"one-lab/internal/handler"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg config.Config, handler handler.Handler) *http.Server {
	mux := http.NewServeMux()
	handler.InitRoutes(mux)
	return &http.Server{
		Addr:           cfg.ServerAddress,
		Handler:        mux,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
}
