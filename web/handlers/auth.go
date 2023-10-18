package handlers

import "github.com/gofiber/fiber/v2"

// Used as a Handler on the Auth middleware returns the 401 Page
func NotAuthorized(c *fiber.Ctx) error {
	return c.Render("errors/401", fiber.Map{
		"Title": "401 - Unauthorized",
	}, "layouts/errors")
}
