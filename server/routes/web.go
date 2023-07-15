package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/livghit/go-htmx/server/controllers"
)

// here we will create the Routes with Templates . Webviews

type Router struct{}

func (r *Router) LoadWebRoutes(a *fiber.App) {
	a.Static("/static", "./client/static")

	a.Get("/", controllers.UserIndex)

}

func (r *Router) LoadMiddleware(a *fiber.App) {}
