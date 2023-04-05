package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/DeS313/cloud-disk/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type regJson struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *MyHandler) registration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	if r.Method == http.MethodPost {
		var input regJson
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			log.Println("Ошибка дешифроки Json")
			return
		}

		u := models.User{Email: input.Email, Password: input.Password}

		user, err := h.service.Create(r.Context(), &u)

		if err != nil {

			if mongo.IsDuplicateKeyError(err) {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("пользователь с таким email уже существует"))
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		res, _ := json.Marshal(user)
		w.Write([]byte(res))
		return

	}

}

func (h *MyHandler) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		jsonRes, err := json.Marshal(map[string]string{"message": "Server error"})
		if err != nil {
			log.Println("Ошибка создание JSON")
			return
		}
		w.Write([]byte(jsonRes))
		return
	}

	var input regJson
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println("Ошибка дешифроки Json")
		return
	}

	u := models.User{Email: input.Email, Password: input.Password}
	u, err := h.service.FindOneByEmail(r.Context(), &u)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		if errors.Is(err, mongo.ErrNoDocuments) {
			w.Write([]byte("пользователь не найден"))
			return
		}
		w.Write([]byte("неверный логин или пароль"))
		return
	}
	token, err := h.service.GenerateToken(u.ID.String())
	if err != nil {
		fmt.Println(err, "handler error")
	}

	reqJson, _ := json.Marshal(map[string]interface{}{
		"token": token,
		"user":  u,
	})
	w.Write(reqJson)
}
