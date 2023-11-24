package main

import (
	"backend/pkg/jsonHelper"
	"backend/pkg/logger"
	"backend/services/userService/config"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	if _, configError := config.SetupConfig(); configError != nil {
		logger.Fatal(configError.Error())
	}

	config.ServerConfig()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response := struct {
			Status  string
			Message string
		}{Status: "alive", Message: "Hello World!"}
		jsonHelper.HttpResponse(&response, w)
	})

	http.ListenAndServe(":5000", r)
}
