package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/DeS313/cloud-disk/internal/config"
	"github.com/DeS313/cloud-disk/internal/db/mongodb"
	"github.com/DeS313/cloud-disk/internal/handlers"
	"github.com/DeS313/cloud-disk/internal/service"
	"github.com/DeS313/cloud-disk/internal/storage"
)

func main() {
	config, err := config.GetConfig("./config/default.json")
	if err != nil {
		fmt.Println(err)
	}

	log.Println("Подключение к базе данных")
	db, err := mongodb.NewClient(context.TODO(), config.DB)
	if err != nil {
		panic(err)
	}
	storage := storage.NewStorage(db, "users")

	service := service.NewService(storage)

	var bindIP = fmt.Sprintf("%v:%v", config.HOST, config.PORT)
	handler := handlers.NewMyHandler(service)
	c := handlers.CorsSetting()

	log.Printf("start web-server on %v", bindIP)

	err = http.ListenAndServe(bindIP, c.Handler(handler.Register()))
	log.Fatal(err)

}
