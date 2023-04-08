package handler

import (
	"net/http"
	"one-lab/internal/service"
)
// handler - это твои контроллеры для определенного транспорта, поэтому их нужно хранить в transport/http/handler.go
type Handler struct {
	services *service.UserService
}

func NewHandler(services *service.UserService) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux() // если у тебя функция отвечает за создание роутов, не нужно здесь создавать твой сервер
	mux.HandleFunc("/add", h.addUser)
	return mux
}
