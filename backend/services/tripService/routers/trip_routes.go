package routers

import (
	"backend/pkg/database"
	"backend/pkg/events"
	"backend/pkg/logger"
	"backend/services/tripService/controller"
	"backend/services/tripService/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func TripRoutes(router chi.Router, db *database.DB) {

	// Create dependencies for trip controller layer
	repo := repository.NewGormRepository(db)
	validator := validator.New(validator.WithRequiredStructEnabled())
	natsConnector, err := events.NewNatsConnector("nats1")
	if err != nil {
		logger.Error(err.Error())
	}

	tripController := controller.NewTripHandler(repo, validator, natsConnector)

	router.Route("/trip", func(r chi.Router) {
		r.Post("/", tripController.CreateTrip)
		r.Get("/{id}", tripController.GetTrip)
		r.Put("/{id}", tripController.UpdateTrip)
		r.Delete("/{id}", tripController.DeleteTrip)
	})

	router.Route("/trips", func(r chi.Router) {
		r.Get("/", tripController.GetTrips)
	})
}
