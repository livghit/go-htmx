package handlers

import "github.com/gofiber/fiber/v2"

func NotAuthorized(c *fiber.Ctx) error {
	return c.Render("errors/401", fiber.Map{
		"Title": "401 - Unauthorized",
	}, "layouts/errors")
}
