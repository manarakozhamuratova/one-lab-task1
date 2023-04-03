package handler

import (
	"net/http"
	"one-lab/internal/service"
)

type Handler struct {
	services *service.UserService
}

func NewHandler(services *service.UserService) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/add", h.addUser)
	return mux
}
