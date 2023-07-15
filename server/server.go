package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/livghit/go-htmx/server/routes"
)

type Server struct {
	name   string
	port   string
	engine *fiber.App
}


func Run() {
	app := fiber.New() 

	router := routes.Router{}
	//if there is any middleware use:
	// router.LoadMiddleware(app)
	router.LoadWebRoutes(app)

  app.Listen(":3000")
}
