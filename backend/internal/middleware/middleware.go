package middleware

import (
	"github.com/DeS313/cloud-disk/internal/service"
)

type Middleware struct {
	service *service.Service
}

// func (m *Middleware) Middleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		if r.Method == http.MethodOptions {
// 			next.ServeHTTP(w, r)
// 		}
// 		w.Header().Set("Content-Type", "application/json")

// 		if err != nil {
// 			log.Println(err, "MIDDLEWARE ERROR")
// 			return
// 		}
// 		r.Body.Read([]byte(tokenDecoded))
// 		next.ServeHTTP(w, r)
// 		return
// 	})
// }

func NewMiddleware(service *service.Service) *Middleware {
	return &Middleware{
		service: service,
	}
}
