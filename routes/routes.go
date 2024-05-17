package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

// defining the Web routes
func Web(r *chi.Mux) {
	// here the routes that you want to have as you're WEBROUTES
}

// defining the Api routes
func Api(r *chi.Mux) {
	// here the routes that you want to have as you're APIROUITES
	// an small best practice is using grouping ex below
	apiRouter := chi.NewRouter()
	apiRouter.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Mount this using r."))
	})

	r.Mount("api", apiRouter)
}
