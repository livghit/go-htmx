package main

import (
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi"
	"github.com/livghit/go-htmx/config"
	"github.com/livghit/go-htmx/routes"
)

func init() {
	config.LoadEnv()
	config.ConnectToDatabase()
  config.StaticFiles("static")
}

func main() {
  	router := chi.NewRouter()
	router.Mount("/", routes.WebRoutes())
	router.Mount("/api/v1", routes.V1Api())

	log.Info("Welcome to " + os.Getenv("APP_NAME"))
	http.ListenAndServe(":3000", router)
}
