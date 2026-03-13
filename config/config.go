package config

import (
	"os"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

// Env holds all application configuration loaded from environment variables.
type Env struct {
	AppName     string
	Port        string
	DBEngine    string // "sqlite" | "postgres" | "mysql"
	DBName      string // sqlite filename or postgres db name
	DBConn      string // full DSN for postgres/mysql
	JWTSecret   string
	JWTExpiry   int    // hours until token expires (default 24)
	FlashSecret string // HMAC key for signing flash cookies
}

// LoadEnv reads the .env file and populates environment variables.
// Call this once at startup before accessing any env values.
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Warn("no .env file found, reading from environment")
	}
}

// Get returns the current Env populated from os environment.
// Applies sensible defaults for development.
func Get() Env {
	expiry, _ := strconv.Atoi(getEnvOr("JWT_EXPIRY", "24"))

	return Env{
		AppName:     getEnvOr("APP_NAME", "go-htmx"),
		Port:        getEnvOr("PORT", ":3000"),
		DBEngine:    getEnvOr("DB_ENGINE", "sqlite"),
		DBName:      getEnvOr("DB_NAME", "app.db"),
		DBConn:      os.Getenv("DB_CONN"),
		JWTSecret:   getEnvOr("JWT_SECRET", "dev-secret-change-in-production"),
		JWTExpiry:   expiry,
		FlashSecret: getEnvOr("FLASH_SECRET", "dev-flash-secret-change-in-production"),
	}
}

func getEnvOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
