package handlers

import "github.com/gofiber/fiber/v2"

func UserHandler(c *fiber.Ctx) error {
	return c.Render("pages/users", fiber.Map{
		"Title": "SIUUU",
	}, "layouts/base")
}
