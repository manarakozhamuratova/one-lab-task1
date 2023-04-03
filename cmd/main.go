package main

import (
	"context"
	"log"
	"one-lab/config"
	"one-lab/internal/handler"
	"one-lab/internal/service"
	"one-lab/internal/storage/memory"
	"one-lab/transport/httpserver"
	"os"
	"os/signal"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := run(); err != nil {
		log.Printf("Error from:%v", err)
		return
	}
}

func run() error {
	config, err := config.LoadConfig(".")
	if err != nil {
		logrus.Fatalf("cannot load config:%v", err)
	}
	storage := memory.NewStorage()
	service := service.NewUserService(storage)
	handler := handler.NewHandler(service)
	srv := new(httpserver.Server)
	go func() {
		if err := srv.NewServer(config, handler.InitRoutes()); err != nil {
			logrus.Fatalf("Error while running server %s", err)
			return
		}
	}()

	logrus.Info("App starting...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	logrus.Info("App shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = srv.Shutdown(ctx); err != nil {
		logrus.Error("error while server shutting down %s", err)
		return err
	}
	return nil
}
