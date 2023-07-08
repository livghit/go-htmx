package server

import (
	"github.com/labstack/echo/v4"
  "github.com/livghit/go-htmx/server/templates"
	"html/template"
)

func Run() {
	t := &Template{
		templates: template.Must(template.ParseGlob("./client/**/*.html")),
	}

	e := echo.New()
	e.Renderer = t
	e.Static("/css", "./client/static/css")
	e.GET("/404", templates.NotFound)
	e.GET("/", templates.Homepage)
	e.Logger.Fatal(e.Start(":1323"))
}
