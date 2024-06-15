package routes

import (


	"github.com/PranitRout07/Practice-Golang/chi_and_htmx/handlers"
	"github.com/PranitRout07/Practice-Golang/chi_and_htmx/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() *chi.Mux {
	routes := chi.NewMux()
	routes.Use(middleware.Logger)

	routes.Get("/", handlers.HomeHandler)
	routes.Get("/posts", handlers.PostHandler)
	routes.Get("/products", handlers.ProductHandler)

	routes.Route("/details/{id}", func(routes chi.Router) {
		routes.Use(middlewares.PostCtx)
		routes.Get("/", handlers.DetailHandler)
		routes.Delete("/delete",handlers.DeleteArticles)
	})

	routes.Post("/addposts",handlers.PostArticles)

	return routes
}
