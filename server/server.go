package server

import (
	"html/template"

	"github.com/labstack/echo/v4"
	"github.com/livghit/go-htmx/server/api"
	"github.com/livghit/go-htmx/server/templates"
)

type Server struct {
	engine *echo.Echo
}

// Initiating the Server Struct with the ECHO Engine
func (s *Server) New() *Server {
	s.engine = echo.New()
	return s
}

func Run() {
	t := &Template{
		templates: template.Must(template.ParseGlob("./client/**/*.html")),
	}
	e := Server{}
	e.New()
	e.engine.Renderer = t
	e.engine.Static("/css", "./client/static/css")
	e.engine.GET("/api/:name", api.HandleSearch)
	e.engine.GET("/404", templates.NotFound)
	e.engine.GET("/", templates.Homepage)
	e.engine.Logger.Fatal(e.engine.Start(":1323"))
}
