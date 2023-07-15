package routes

import "github.com/gofiber/fiber/v2"

func (r *Router) LoadApiRoutes(a *fiber.App) {
	a.Static("/static", "./client/static")

}
