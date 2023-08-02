package handlers 

import "github.com/gofiber/fiber/v2"

/*
 Templates are here to represent the Handlers for each View (Like controllers in an MVC app)
*/

func HomepageHandler(c *fiber.Ctx) error {

	return c.SendString("hello homepage")
}
