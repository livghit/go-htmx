package storage

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/mssql/v2"
	"github.com/gofiber/storage/postgres/v3"
	"github.com/gofiber/storage/sqlite3/v2"
	"github.com/livghit/go-htmx/pkg/config"
)

func New(env config.Env) fiber.Storage {
	var storage fiber.Storage

	switch env.DB_ENGINE {
	case "mysql":
		storage = mssql.New(config.StorageMysqlConfig())
		return storage
	case "sqlite":
		storage = sqlite3.New(config.StorageSqliteConfig())
		return storage
	case "postgress":
		storage = postgres.New(config.StoragePostgressConfig())
		return storage
	default:
		panic("No supported Engine selected for Storage")
	}
}
