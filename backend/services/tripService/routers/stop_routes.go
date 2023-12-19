package routers

import (
	"backend/pkg/database"
	"backend/services/tripService/controller"
	"backend/services/tripService/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func StopRoutes(router chi.Router, db *database.DB) {
	repo := repository.NewGormRepository(db)
	validator := validator.New(validator.WithRequiredStructEnabled())
	stopController := controller.NewStopHandler(repo, validator)

	router.Route("/stop", func(r chi.Router) {
		r.Post("/", stopController.CreateStop)
	})

	router.Route("/stops", func(r chi.Router) {
		r.Post("/", stopController.CreateStops)
	})
}
