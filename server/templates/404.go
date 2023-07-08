package templates

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func NotFound(c echo.Context) error {
	return c.Render(http.StatusOK, "404", "")
}
