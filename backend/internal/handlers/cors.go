package handlers

import (
	"net/http"

	"github.com/rs/cors"
)

func CorsSetting() *cors.Cors {
	cors := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodPatch,
			http.MethodOptions,
		},
		AllowedOrigins: []string{
			"http://localhost:3000",
			"http://localhots:3000/",
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"Authorization",
			"Content-Type",
		},
		OptionsPassthrough: true,
		ExposedHeaders: []string{
			"Content-Type",
			"id",
		},
		Debug: true,
	})

	return cors
}
