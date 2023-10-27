package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/livghit/go-htmx/pkg/config"
	"github.com/livghit/go-htmx/pkg/data"
)

func main() {
	// this loads the env files
	env, err := config.LoadEnv(".")
	if err != nil {
		log.Printf("%v", err)
	}
	log.Printf("%v lock and loaded !", env.APPNAME)

	if err := run(env); err != nil {
		log.Fatal(err)
	}

}

func run(env config.Env) error {
	app := fiber.New(config.ViewConfigs())
	//initiating the database and connecting to the engine set up inside the env file
	database := data.InitDatabase()
	database.Connect(env)
  database.MakeMigration()
  
	//Here you can register own routes
	SetupApiRoutes(app)

	return app.Listen(env.PORT)

}
