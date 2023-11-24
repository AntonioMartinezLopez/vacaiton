package main

import (
	"backend/pkg/database"
	"backend/pkg/jsonHelper"
	"backend/pkg/logger"
	"backend/services/userService/config"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	// Check configuration and for all env variables
	if _, configError := config.SetupConfig(); configError != nil {
		logger.Fatal(configError.Error())
	}

	// Get server configuration
	config.ServerConfig()

	// Generate data source name for database connection
	postgresDsn := config.GetDSNConfig()

	// create database connection and watch for connection
	db, err := database.DBConnection(postgresDsn)
	if err != nil {
		logger.Fatal("%v", err)
	}
	go database.WatchDBConnection(db)

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
