package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/livghit/go-htmx/pkg/config"
)

type Server struct {
	*fiber.App
	fiber.Storage
}

func New(storage fiber.Storage) *Server {
	return &Server{
		App:     fiber.New(config.ViewConfigs()),
		Storage: storage,
	}
}
