package middleware

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/livghit/go-htmx/web/handlers"
	"github.com/livghit/go-htmx/web/views/errors"
)

type User struct {
	name string
}

// using decorator pattern to redefine the AUTH logic inside the app
func Auth(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// inside here you chek for the authorisation rules that you want  ex
		// here implement JWT and check for the valid token existence
		// if there can
		user := User{
			name: "dan",
		}
		if user.name != "dan" {
			// here we could also render a Unauth View using the engine !
			// Example with html:
			//c.Render("errors/401", fiber.Map{"Title": "Unautherized"}, "layout/base")
			// Example with templ
			handlers.Render(c, errors.NotAuth())
			return c.SendStatus(http.StatusUnauthorized)
		}
		log.Default().Print("JWT MIDDLEWARE INITIATED !")
		return fn(c)
	}
}
