package server

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Function to be put inside the Template
func Homepage(c echo.Context) error {
	return c.Render(http.StatusOK, "base", "homepage")
}

func NotFound(c echo.Context) error {
	return c.Render(http.StatusOK, "404", "")
}
