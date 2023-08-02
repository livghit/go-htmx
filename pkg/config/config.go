package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/livghit/linkhub/pkg/templating"
	"github.com/livghit/linkhub/web/handlers"
)

// Function to pass the Templating Engine to the fiber app
func ViewConfigs() fiber.Config {
	// maybe do engine instead of mustache and have all the engine under engine.NewDjango NewMustache etc?
	engine := templating.HtmlEngine()
	engine.Reload(true)
	configs := fiber.Config{
		Views: engine,
	}

	return configs
}

func AuthConfig() basicauth.Config {
	authconfig := basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
		Unauthorized: handlers.NotAuthorized,
	}

	return authconfig
}

func CsrfConfig() csrf.Config {
	csrfconfig := csrf.Config{
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     "csrf_",
		CookieSameSite: "Lax",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
	}

	return csrfconfig
}
