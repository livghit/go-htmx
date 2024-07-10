package routes

import (
	"github.com/go-chi/chi"
	"github.com/livghit/go-htmx/services"
)

// defining the Web routes
func WebRoutes() *chi.Mux {
	router := chi.NewRouter()
	// here the routes that you want to have as you're WEBROUTES

	router.Get("GET /", services.HandleHomepage)

	return router
}

// defining the Api routes
func V1Api() *chi.Mux {
	// here the routes that you want to have as you're APIROUITES
	// an small best practice is using grouping ex below
	router := chi.NewRouter()
	router.Get("/users", services.GetAllUsers)

	return router
}
