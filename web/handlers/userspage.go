package handlers

import "github.com/gofiber/fiber/v2"

func UserHandler(c *fiber.Ctx) error {
	return c.Render("pages/user, fiber.Map{
		"Title": "SIUUU",
	}, "layouts/base")
}
