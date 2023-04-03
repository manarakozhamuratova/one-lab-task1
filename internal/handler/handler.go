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

func (h *Handler) InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/add", h.addUser)
}
