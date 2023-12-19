package routers

import (
	"backend/pkg/database"
	"backend/services/tripService/controller"
	"backend/services/tripService/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func TripRoutes(router chi.Router, db *database.DB) {
	repo := repository.NewGormRepository(db)
	validator := validator.New(validator.WithRequiredStructEnabled())
	tripController := controller.NewTripHandler(repo, validator)

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
