package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/livghit/linkhub/pkg/config"
	"github.com/livghit/linkhub/pkg/middleware"

	"github.com/livghit/linkhub/web/handlers"
)

func main() {

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
	//Here you can register own routes
	app.Get("/", middleware.Auth(handlers.HomepageHandler))
	app.Get("/users", handlers.UserHandler)

	return app.Listen(env.PORT)

}
