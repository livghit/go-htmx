package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/livghit/go-htmx/routes"
)

func main() {
	router := chi.NewRouter()
	router.Mount("", routes.WebRoutes())
	router.Mount("/api", routes.ApiRoutes())

	http.ListenAndServe(":3000", router)
}
