package middleware

import (
	"net/http"
	"strings"

	"github.com/DeS313/cloud-disk/internal/service"
)

func AuthMiddleware(hand func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if len(strings.Split(r.Header.Get("Authorization"), " ")) <= 1 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Auth error"))
			return
		}
		token := strings.Split(r.Header.Get("Authorization"), " ")[1]
		id, err := service.ParseToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Auth error"))
			return
		}
		r.Header.Add("id", id)
		hand(w, r)

	}
}
