package controllers

import "github.com/gofiber/fiber/v2"

// controler for the User Model

func UserIndex(c *fiber.Ctx) error {
	return c.SendString("I'm a GET request!")
}
