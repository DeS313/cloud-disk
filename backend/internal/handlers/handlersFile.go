package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/DeS313/cloud-disk/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *MyHandler) FileHandleFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet || r.Method == http.MethodOptions {
		h.GetFile(w, r)
		return
	}
	if r.Method == http.MethodPost || r.Method == http.MethodOptions {
		h.CreateFile(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (h *MyHandler) GetFile(w http.ResponseWriter, r *http.Request) {
	var input *models.Files
	// defer r.Body.Close()
	// if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("фа"))
	// 	return

	// }
	oid, err := primitive.ObjectIDFromHex(r.Header.Get("id"))
	if err != nil {
		log.Println(err)
		return
	}
	input.UserID = oid
	file, err := h.service.FindFile(r.Context(), input)
	if err != nil {
		log.Println(err)
		return
	}
	res, _ := json.Marshal(map[string]interface{}{
		"file": file,
	})
	w.Write([]byte(res))
}

func (h *MyHandler) CreateFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input *models.CreateFile
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err)
		return
	}
	oid, _ := primitive.ObjectIDFromHex(r.Header.Get("id"))
	id, err := h.service.CreateFile(r.Context(), &models.Files{
		UserID:    oid,
		Name:      input.Name,
		ParrentID: input.Parent,
		Type:      input.Type,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ошибка создания файла"))
		log.Println(err)
		return
	}

	w.Write([]byte(id))
}
