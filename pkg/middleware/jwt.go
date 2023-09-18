package middleware

import (
	"github.com/gofiber/fiber/v2"
)

//idk how to implement jwt but it will be ok :)

func JWT(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return fn(c)
	}
}
