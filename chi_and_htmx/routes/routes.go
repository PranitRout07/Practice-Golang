package routes

import (
	"github.com/PranitRout07/Practice-Golang/chi_and_htmx/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() *chi.Mux {
	routes := chi.NewMux()
	routes.Use(middleware.Logger)

	routes.Get("/", handlers.HomeHandler)
	routes.Get("/posts", handlers.PostHandler)
	routes.Get("/products", handlers.ProductHandler)

	return routes
}
