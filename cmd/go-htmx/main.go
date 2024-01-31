package main

import (
	"log"

	"github.com/livghit/go-htmx/pkg/config"
	"github.com/livghit/go-htmx/pkg/server"
)

func main() {
	// this loads the env files
	env, err := config.LoadEnv(".")
	if err != nil {
		log.Printf("%v", err)
	}
	log.Printf("%v lock and loaded !", env.APP_NAME)

	if err := run(env); err != nil {
		log.Fatal(err)
	}
}

func run(env config.Env) error {
	app := server.New()
	// Here you can register own routes
	SetupApiRoutes(app)

	return app.Listen(env.APP_PORT)
}
