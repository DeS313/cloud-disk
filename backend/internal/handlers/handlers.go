package handlers

import (
	"net/http"

	"github.com/DeS313/cloud-disk/internal/middleware"
	"github.com/DeS313/cloud-disk/internal/service"
)

type MyHandler struct {
	handler *http.ServeMux
	service *service.Service
}

const (
	REGISTRATION = "/registration"
	LOGIN        = "/login"
	USER         = "/user"
	FILES        = "/files"
)

func (h *MyHandler) Register() *http.ServeMux {
	h.handler.HandleFunc(REGISTRATION, h.registration)
	h.handler.HandleFunc(LOGIN, h.login)
	h.handler.HandleFunc(USER, middleware.AuthMiddleware(h.getUser))
	h.handler.HandleFunc(FILES, middleware.AuthMiddleware(h.FileHandleFunc))
	h.handler.HandleFunc("/upload", middleware.AuthMiddleware(h.UploadFile))
	return h.handler
}

func NewMyHandler(service *service.Service) *MyHandler {
	return &MyHandler{
		handler: http.NewServeMux(),
		service: service,
	}
}
