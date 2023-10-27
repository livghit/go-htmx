package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/livghit/go-htmx/pkg/templating"
	"github.com/livghit/go-htmx/web/handlers"
	"github.com/spf13/viper"
)

type Env struct {
	PORT    string `mapstructure:"PORT"`
	APPNAME string `mapstructure:"APP_NAME"`
	DBENGINE string `mapstructure:"DB_ENGINE"`
  DBNAME string `mapstructure:"DB_NAME"`

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
		Unauthorized: handlers.NotAuthorized,
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
