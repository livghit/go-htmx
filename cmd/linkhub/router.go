package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/livghit/linkhub/web/handlers"
)

func SetupWebRoutes(app *fiber.App) {
	// Setting up the routes for the web
	// this means inside here you map the url with an handler ex
	app.Get("/", handlers.HomepageHandler)

}

func SetupApiRoutes(app *fiber.App) {
	// Here you setup the ROUTE GROUP for the Api
	api := app.Group("/api/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "1")
		return c.Next()
	})

	// on the ROUTE GROUP you now can register the routes
	// you can create api handler in the handlers dir and add handler.HandlerName to the get instead of the func
	api.Get("/path", func(c *fiber.Ctx) error {
		return c.SendString("This matches the path /api/v1/path")
	})

}
