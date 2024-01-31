package main

import (
	"log"

	"github.com/livghit/go-htmx/pkg/config"
	"github.com/livghit/go-htmx/pkg/server"
	"github.com/livghit/go-htmx/pkg/storage"
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
	s := storage.New(env)
	app := server.New(s)

	// Here you can register own routes example below:
	SetupApiRoutes(app.App)
	SetupWebRoutes(app.App)

	return app.Listen(env.APP_PORT)
}
