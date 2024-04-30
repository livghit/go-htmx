package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/livghit/go-htmx/pkg/config"
)

// In this dir we setup the middleware provided by fiber itself
// the config for it will be found in config.go

// Adding the auth middleware to the application
//
//	 func Auth() func(c *fiber.Ctx) error {
//		auth := basicauth.New(config.AuthConfig())
//		return auth
//	}
//

// Adding the csrf middleware to the application
func CSRF() func(c *fiber.Ctx) error {
	csrf := csrf.New(config.CsrfConfig())
	return csrf
}
