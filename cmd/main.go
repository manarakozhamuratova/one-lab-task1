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
)

func main() {
	if err := run(); err != nil {
		log.Printf("Error from:%v", err)
		return
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	shutdown(cancel)
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	storage := memory.NewStorage(ctx)
	service := service.NewUserService(storage)
	handler := handler.NewHandler(service)
	srv := httpserver.NewServer(config, *handler)
	log.Println("http://localhost:9090/")
	return srv.ListenAndServe()
}

func shutdown(ctx context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		log.Print(<-c)
		ctx()
	}()
}
