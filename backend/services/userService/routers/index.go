package routers

import (
	"backend/pkg/database"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router *chi.Mux, db *database.DB) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("\"live\": \"ok\""))
	})
	//Add All route
	//ExamplesRoutes(router, db)
}
