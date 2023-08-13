package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/livghit/linkhub/pkg/config"
	"github.com/livghit/linkhub/pkg/middleware"

	"github.com/livghit/linkhub/web/handlers"
)


func main() {

	app := fiber.New(config.ViewConfigs())
	env, err := config.LoadEnv(".")
	if err != nil {
		log.Printf("%v", err)
	}
	log.Printf("%v lock and loaded !", env.APPNAME)

	//Here you can register own routes  
	app.Get("/", middleware.Auth(handlers.HomepageHandler))
	app.Get("/users", handlers.UserHandler)

	log.Fatal(app.Listen(env.PORT))
}
