package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/livghit/linkhub/pkg/config"
)

// Adding the auth middleware to the application
func Auth() func(c *fiber.Ctx) error {
	auth := basicauth.New(config.AuthConfig())

	return auth
}

// Adding the csrf middleware to the application
func CSRF() func(c *fiber.Ctx) error {
	csrf := csrf.New(config.CsrfConfig())

	return csrf
}
