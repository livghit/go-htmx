package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/deez", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Nuts!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}

