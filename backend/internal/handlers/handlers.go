package handlers

import (
	"net/http"

	"github.com/DeS313/cloud-disk/internal/service"
)

type MyHandler struct {
	handler *http.ServeMux
	service *service.Service
}

const (
	REGISTRATION = "/registration"
)

func (h *MyHandler) Register() *http.ServeMux {
	h.handler.HandleFunc(REGISTRATION, h.registration)

	return h.handler
}

func NewMyHandler(service *service.Service) *MyHandler {
	return &MyHandler{
		handler: http.NewServeMux(),
		service: service,
	}
}
