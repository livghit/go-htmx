package main

import (
	"log"

	"github.com/livghit/go-htmx/config"
	"github.com/livghit/go-htmx/db"
	"github.com/livghit/go-htmx/db/models"
	"golang.org/x/crypto/bcrypt"
)

// seed creates a default admin user for development.
// Run with: make seed
func main() {
	config.LoadEnv()
	cfg := config.Get()

	if err := db.Connect(cfg.DBEngine, cfg.DBName); err != nil {
		log.Fatalf("connect: %v", err)
	}
	defer db.Close()

	if err := db.Migrate(); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("hash: %v", err)
	}

	user, err := models.Create("admin", "admin@example.com", string(hash))
	if err != nil {
		log.Fatalf("create user: %v", err)
	}

	log.Printf("seeded user: id=%d username=%s email=%s", user.ID, user.Username, user.Email)
}
