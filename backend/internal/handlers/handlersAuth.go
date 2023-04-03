package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type regJson struct {
	Email    string `json:"email"`
	Password string `json:"passowrd"`
}

func (h *MyHandler) registration(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		var input regJson
		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			log.Fatalln("Ошибка дешифроки Json")
			return
		}

	}

	log.Fatal(http.StatusMethodNotAllowed)
	w.WriteHeader(http.StatusMethodNotAllowed)

	jsonRes, err := json.Marshal(map[string]string{"message": "Server error"})
	if err != nil {
		log.Fatal("Ошибка создание JSON")
		return
	}
	w.Write([]byte(jsonRes))
}
