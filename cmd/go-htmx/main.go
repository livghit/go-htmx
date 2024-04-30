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

// Runner function witch starts the server and loads all the needed things
func run(env config.Env) error {
	s := storage.New(env)
	app := server.New(s)

	/*  Here you can register own routes example below:
	* 	we pass the storage here so we can use it
	* 	inside the handlers to modify and store
	* 	the data
	 */
	ApiRoutes(app.App)
	WebRoutes(app.App)

	return app.Listen(env.APP_PORT)
}
