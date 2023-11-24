package routers

import (
	"backend/pkg/database"
	"backend/services/userService/controller"
	"backend/services/userService/repository"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(router *chi.Mux, db *database.DB) {
	repo := repository.NewGormRepository(db)
	userController := controller.NewUserHandler(repo)
	router.Group(func(r chi.Router) {
		r.Post("/create", userController.CreateUser)

	})
}
