package routers

import (
	"backend/pkg/database"
	"backend/services/userService/controller"
	"backend/services/userService/repository"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(router chi.Router, db *database.DB) {
	repo := repository.NewGormRepository(db)
	userController := controller.NewUserHandler(repo)

	router.Route("/auth", func(r chi.Router) {
		r.Post("/signup", userController.CreateUser)
		r.Get("/{user_id}", userController.GetUserInfo)
		r.Get("/login", userController.LoginUser)
		r.Get("/logout", userController.LogoutUser)
		r.Get("/", userController.CheckTokenValid)
	})
}
