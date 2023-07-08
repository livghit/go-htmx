package templates 

import (
  "net/http"
	"github.com/labstack/echo/v4"
)

// Function to be put inside the Template
func Homepage(c echo.Context) error {
	return c.Render(http.StatusOK, "Homepage", "homepage")
}
