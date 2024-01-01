package main

import (

	// docs is generated by Swag CLI, you have to import it.
	"backend/pkg/database"
	"backend/pkg/logger"
	"backend/services/userService/config"
	"backend/services/userService/migrations"
	"backend/services/userService/routers"
	"net/http"
	"time"

	docs "backend/docs/userService"

	"github.com/spf13/viper"
)

//	@title			User/Auth API
//	@version		1.0
//	@description	This server is used for creating new users and conduct authentication
//	@descripition	It is possible to sign up using email and password or using the oauth google client
//	@description	Swagger authentication is set to oauth to make login easier
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Vacaition API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@Tags	Auth, OAuth

//	@BasePath	/userservice/api

//	@securitydefinitions.oauth2.implicit	OAuth2Application
//	@authorizationurl						http://localhost:8080/userservice/api/oauth?provider=google
//	@scope.write							Grants write access
//	@scope.admin							Grants read and write access to administrative information

func main() {

	// set swagger
	docs.SwaggerInfo.Host = viper.GetString("URL")

	// Check configuration and for all env variables
	if _, configError := config.SetupConfig(); configError != nil {
		logger.Fatal(configError.Error())
	}

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
