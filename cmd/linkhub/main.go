package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/livghit/linkhub/pkg/config"
	"github.com/livghit/linkhub/pkg/middleware"

	// "github.com/livghit/linkhub/pkg/middleware"
	"github.com/livghit/linkhub/web/handlers"
)

/*
QUESTIONS :
- Where will the middleware be placed
- How to handle user login with azur ad
*/

func main() {

	app := fiber.New(config.ViewConfigs())
	env, err := config.LoadEnv(".")
	if err != nil {
		log.Printf("%v", err)
	}
	log.Printf("%v", env.APPNAME)

	// From here register all routes
	app.Get("/", middleware.Auth(handlers.HomepageHandler))

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"Title": "App",
		}, "layouts/base")
	})

	app.Get("/users", handlers.UserHandler)

	log.Fatal(app.Listen(env.PORT))

}
