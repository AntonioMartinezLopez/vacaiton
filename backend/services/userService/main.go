package main

import (
	"backend/pkg/database"
	"backend/pkg/logger"
	"backend/services/userService/config"
	"backend/services/userService/migrations"
	"backend/services/userService/routers"
	"net/http"
	"time"
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

	migrations.Migrate(db)

	// Initialize Router
	router := routers.SetupRouter(db)
	server := http.Server{
		Addr:              config.ServerConfig(),
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}
	logger.Fatal("%v", server.ListenAndServe())
}
