package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/livghit/go-htmx/routes"
)

func main() {
	router := chi.NewRouter()

	routes.Web(router)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello something"))
	})

	http.ListenAndServe(":3000", router)
}
