package main

import (

	// docs is generated by Swag CLI, you have to import it.
	"backend/pkg/database"
	"backend/pkg/events"
	"backend/pkg/logger"
	"backend/services/tripService/config"
	"backend/services/tripService/migrations"
	"backend/services/tripService/routers"
	"fmt"
	"net/http"
	"time"

	docs "backend/docs/tripService"

	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
)

//	@title			Trip API
//	@version		1.0
//	@description	This server is used for creating new trips
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Vacaition API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@Tags	Trip, Stop

//	@BasePath	/tripservice/api

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

	// test
	events.ExampleJetStreamGroupConsumer("FOO", "foo", "testgroup", func(m *nats.Msg) {
		logger.Log("Message: " + string(m.Data))
		meta, _ := m.Metadata()

		logger.Log(fmt.Sprintf("Stream Sequence  : %v\n", meta.Sequence.Stream))
		logger.Log(fmt.Sprintf("Consumer Sequence: %v\n", meta.Sequence.Consumer))
	})

	natsConnector, err := events.NewNatsConnector("nats1")
	if err != nil {
		logger.Error(err.Error())
	}

	natsConnector.PublishStream("foo", []byte("Hello JS Async!"))

	// Publish messages asynchronously.
	for i := 0; i < 10; i++ {
		natsConnector.PublishStreamAsync("foo", []byte("Hello JS Async!"))
	}
	select {
	case <-natsConnector.PublishAsyncFinished():
	case <-time.After(5 * time.Second):
		fmt.Println("Did not resolve in time")
	}
	logger.Fatal("%v", server.ListenAndServe())
}
