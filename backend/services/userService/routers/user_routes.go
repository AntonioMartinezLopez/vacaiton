package routers

import (
	"backend/pkg/database"
	"backend/services/userService/controller"
	"backend/services/userService/middleware"
	"backend/services/userService/repository"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(router chi.Router, db *database.DB) {
	repo := repository.NewGormRepository(db)
	userController := controller.NewUserHandler(repo)

	router.Route("/auth", func(r chi.Router) {
		r.Post("/signup", userController.CreateUser)
		r.Post("/login", userController.LoginUser)

		r.Group(func(r chi.Router) {
			r.Use(middleware.JwtGuard)
			r.Get("/logout", userController.LogoutUser)
			r.Get("/", userController.CheckTokenValid)
			r.Get("/user", userController.GetUserInfo)
		})
	})
}
