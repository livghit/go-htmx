package routes

import (
	"github.com/go-chi/chi"
	"github.com/livghit/go-htmx/services"
)

// defining the Web routes
func WebRoutes() *chi.Mux {
	webRoutes := chi.NewRouter()
	// here the routes that you want to have as you're WEBROUTES

	return webRoutes
}

// defining the Api routes
func ApiRoutes() *chi.Mux {
	// here the routes that you want to have as you're APIROUITES
	// an small best practice is using grouping ex below
	apiRouter := chi.NewRouter()
	apiRouter.Get("/users", services.GetAllUsers)

	return apiRouter
}
