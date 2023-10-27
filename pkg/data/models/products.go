package models

import (
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	Id          int32
	Name        string
	Description string
}

func MigrateProductTable() string {
	return `CREATE TABLE IF NOT EXISTS applications(
		"idApplication" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"description" TEXT		
	  );`
}

func GetProducts(ctx *fiber.Ctx) error {
	// returning all the products
	return ctx.JSON("Products")
}

func GetProductById() {
	// returning product by id
}

// Easy Way of creating a Product
func Create() {
	//
}

func Update() {
	// retrning an updated product after that saving it to the db
}

func Delete() {
}
