package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func HandleSearch(c echo.Context) error {
	time.Sleep(time.Second * 5)
	name := c.Param("name")
	return c.JSON(http.StatusOK, name)
}
