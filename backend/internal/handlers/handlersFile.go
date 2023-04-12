package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/DeS313/cloud-disk/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *MyHandler) FileHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodOptions {
		return
	}
	if r.Method == http.MethodGet {
		log.Println("dfkdfslkdfslkfsdlkfsdlkdfs")
		h.GetFile(w, r)
		return
	}
	if r.Method == http.MethodPost {
		h.CreateFile(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (h *MyHandler) GetFile(w http.ResponseWriter, r *http.Request) {
	parant := r.URL.Query().Get("parant")
	file, err := h.service.FindFile(r.Context(), r.Header.Get("id"), parant)
	if err != nil {
		log.Println(err, "SERVICE ERROR")
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
	}
	file, err := h.service.FindOneFile(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ошибка создания файла"))
		log.Println(err)
		return
	}
	res, _ := json.Marshal(file)
	w.Write([]byte(res))
}

type input struct {
	Parrent string `json:"parent"`
	Type    string `json:"file"`
}

func (h *MyHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	// 10<<20 10mb
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Println(err, 1)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// defer r.Body.Close()
	// if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
	// 	log.Println(err, 2)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	input := input{
		Parrent: r.FormValue("parent"),
		Type:    r.FormValue("type"),
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println(err, 3)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %-v\n", handler.Header)

	user, err := h.service.FindOne(r.Context(), r.Header.Get("id"))
	if err != nil {
		log.Println(err, 4)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user.UserSpace+int(handler.Size) > user.DiskSpace {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(user.UserSpace+int(handler.Size) > user.DiskSpace)
		w.Write([]byte(fmt.Sprintf("There no space on the disk. UserSpace:%v, FileSize:%v, DiskSpace:%v", user.UserSpace, handler.Size, user.DiskSpace)))
		return
	}
	user.UserSpace = user.UserSpace + int(handler.Size)
	if err := h.service.Update(r.Context(), &user); err != nil {
		log.Println(err, 5)
		w.Write([]byte(err.Error()))
		return
	}

	var path string

	newFile := &models.Files{
		Name:   handler.Filename,
		Size:   int(handler.Size),
		Type:   "file",
		UserID: user.ID,
	}

	if input.Parrent == "" {
		path = fmt.Sprintf("/home/des/Рабочий стол/my_project/cloud_storage/files/%s%s", user.ID.Hex(), handler.Filename)
		newFile.ParrentID = user.ID
		newFile.Path = path
	} else {
		parent, _ := h.service.FindOneFile(r.Context(), input.Parrent)
		path = fmt.Sprintf("/home/des/Рабочий стол/my_project/cloud_storage/files/%s/%s/%s", user.ID.Hex(), parent.Path, handler.Filename)
		newFile.Path = path
		newFile.ParrentID = parent.ID
	}

	tempFile, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY|os.O_RDONLY, 0777)
	if err != nil {
		log.Println(err, 6)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err, 7)
		w.Write([]byte(err.Error()))
		return
	}

	tempFile.Write(fileBytes)
	fmt.Println(newFile.ParrentID, "parent")

	id, err := h.service.CreateF(r.Context(), newFile)
	if err != nil {
		fmt.Println(err, 8)
		w.Write([]byte(err.Error()))
		return
	}
	f, _ := h.service.FindOneFile(r.Context(), id)
	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(f)
	w.Write([]byte(res))
	log.Println("file success uploaded")
}
