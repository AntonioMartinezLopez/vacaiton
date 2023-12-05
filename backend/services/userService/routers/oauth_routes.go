package routers

import (
	"backend/pkg/database"
	"backend/services/userService/controller"
	"backend/services/userService/repository"

	"github.com/go-chi/chi/v5"
)

func OauthRoutes(router chi.Router, db *database.DB) {
	repo := repository.NewGormRepository(db)
	oauthController := controller.NewOauthHandler(repo)

	router.Route("/oauth", func(r chi.Router) {
		r.Get("/callback", oauthController.OauthCallback)
		r.Get("/logout", oauthController.OauthLogout)
		r.Get("/", oauthController.OauthInit)
	})
}
