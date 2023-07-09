package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleSearch(c echo.Context) error {
	name := c.Param("name")
	return c.JSON(http.StatusOK, name)
}
