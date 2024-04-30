package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/livghit/go-htmx/web/views/pages"
)

// Used as a Handler on the Auth middleware returns the 401 Page
func RenderRegisterPage(c *fiber.Ctx) error {
	return Render(c, pages.RegisterPage())
}

func RenderLoginPage(c *fiber.Ctx) error {
	return Render(c, pages.LoginPage())
}

func Login(c *fiber.Ctx) error {
	log.Fatal(c.Context().PostBody())
	return nil
}
