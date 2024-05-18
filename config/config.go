package config

import (
	"log/slog"

	"github.com/joho/godotenv"
)

type Env struct {
	/**
	* define the fields inside you're env
	* in order to have them ready to use
	* ex. Env.DB_NAME and so on
	 */
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		slog.Error(".env file was not found. Are you sure you have one ?")
	}
}

/**
* This will take the database type definde in the env and also all
* the needed inputs for it ex pw , connection string and so on
 */
func ConnectToDatabase() {
	var connection string = "Hello"
	switch connection {
	}
}
