package server

import (
	"github.com/labstack/echo/v4"
	"html/template"
)

func Run() {
	t := &Template{
		templates: template.Must(template.ParseGlob("./client/*.html")),
	}

	e := echo.New()
	e.Renderer = t
  e.Static("/css" , "./client/static/css")
	e.GET("/", Homepage)
	e.Logger.Fatal(e.Start(":1323"))
}
