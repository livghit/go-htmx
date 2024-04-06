package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/storage/mssql/v2"
	"github.com/gofiber/storage/postgres/v3"
	"github.com/gofiber/storage/sqlite3/v2"
	"github.com/livghit/go-htmx/pkg/templating"
	"github.com/livghit/go-htmx/web/handlers"
	"github.com/spf13/viper"
)

type Env struct {
	APP_PORT   string `mapstructure:"APP_PORT"`
	APP_NAME   string `mapstructure:"APP_NAME"`
	DB_ENGINE  string `mapstructure:"DB_ENGINE"`
	DB_NAME    string `mapstructure:"DB_NAME"`
	JWT_SECRET string `mapstructure:"JWT_SECRET"`
}

func LoadEnv(path string) (env Env, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&env)
	return
}

// Function to pass the Templating Engine to the fiber app
func ViewConfigs() fiber.Config {
	// maybe do engine instead of mustache and have all the engine under engine.NewDjango NewMustache etc?
	engine := templating.HtmlEngine()
	engine.Reload(true)
	configs := fiber.Config{
		Views: engine,
	}

	return configs
}

func AuthConfig() basicauth.Config {
	authconfig := basicauth.Config{
		Users: map[string]string{
			"john":  "doe",
			"admin": "123456",
		},
		Unauthorized: handlers.RenderLoginPage,
	}

	return authconfig
}

func CsrfConfig() csrf.Config {
	csrfconfig := csrf.Config{
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     "csrf_",
		CookieSameSite: "Lax",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
	}

	return csrfconfig
}

/*
* For Documentation on the posible Configs on each Storage please
* take a look at the Fiber Documentation
* */
func StorageSqliteConfig() sqlite3.Config {
	return sqlite3.Config{}
}

func StorageMysqlConfig() mssql.Config {
	return mssql.Config{}
}

func StoragePostgressConfig() postgres.Config {
	return postgres.Config{}
}
