package routers

import (
	"backend/pkg/database"
	"backend/pkg/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(db *database.DB) *chi.Mux {

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middlewares.Cors())
	router.Use(middlewares.UserClaims)

	RegisterRoutes(router, db)
	return router
}
