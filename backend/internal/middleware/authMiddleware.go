package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/DeS313/cloud-disk/internal/service"
)

func AuthMiddleware(hand func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			return
		}
		log.Println(r.Header.Get("Authorization"), "authMiddleware r.get()")
		if len(strings.Split(r.Header.Get("Authorization"), " ")) <= 1 {
			log.Println(r.Header.Get("Authorization"))
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
		r.Header.Del("id")
		return
	}
}
