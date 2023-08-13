package middleware

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type User struct {
	name string
}

// using decorator pattern to redefine the AUTH logic inside the app
func Auth(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// inside here you chek for the authorisation rules that you want  ex
		// I need to implement jwt
		user := User{
			name: "siu",
		}
		if user.name != "dan" {
			// here we could also render a Unauth View using the engine !
			// Example:
			c.Render("errors/401", fiber.Map{"Title": "Unautherized"}, "layout/base")
			return c.SendStatus(http.StatusUnauthorized)
		}

		return fn(c)
	}
}
