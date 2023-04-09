package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/DeS313/cloud-disk/internal/models"
	"github.com/DeS313/cloud-disk/internal/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		oid, _ := primitive.ObjectIDFromHex(user)
		file, err := h.service.CreateFile(r.Context(), &models.Files{
			UserID: oid,
			Name:   "",
		})
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		res, _ := json.Marshal(map[string]interface{}{
			"user_id": user,
			"file_id": file,
		})
		w.Write([]byte(res))
		return

	}

}

func (h *MyHandler) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost && r.Method != http.MethodOptions {
		w.WriteHeader(http.StatusMethodNotAllowed)
		resJson, err := json.Marshal(map[string]string{"message": "Server error"})
		if err != nil {
			log.Println("Ошибка создание JSON")
			return
		}
		w.Write([]byte(resJson))
		return
	}

	var input regJson
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println("Ошика дешифровки JSON")
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
		w.Write([]byte("пользователь логин или пароль"))
		return
	}

	token, err := h.service.GenerateToken(u.ID.Hex())
	if err != nil {
		log.Println(err)
	}

	resJson, _ := json.Marshal(map[string]interface{}{
		"token": token,
		"user":  u,
	})

	w.Write(resJson)

}

func (h *MyHandler) getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("HI")
	if r.Method != http.MethodGet && r.Method != http.MethodOptions {
		w.WriteHeader(http.StatusMethodNotAllowed)
		resJson, err := json.Marshal(map[string]string{"message": "Server error"})
		if err != nil {
			log.Println("Ошибка создание JSON")
			return
		}
		w.Write([]byte(resJson))
		return
	}

	token := strings.Split(r.Header.Get("Authorization"), " ")[1]

	id, err := service.ParseToken(token)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(id)

	user, err := h.service.FindOne(r.Context(), id)
	if err != nil {
		log.Println(err)
		return
	}

	res, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(res)
}
