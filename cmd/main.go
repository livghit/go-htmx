package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/livghit/go-htmx/config"
	"github.com/livghit/go-htmx/db"
	"github.com/livghit/go-htmx/routes"
)

func main() {
	// Load .env file (or read from environment)
	config.LoadEnv()
	cfg := config.Get()

	log.Info("starting", "app", cfg.AppName)

	// Connect to the database and run pending migrations
	if err := db.Connect(cfg.DBEngine, cfg.DBName); err != nil {
		log.Fatal("database connection failed", "err", err)
	}
	defer db.Close()

	if err := db.Migrate(); err != nil {
		log.Fatal("migrations failed", "err", err)
	}

	// Build the router
	router := chi.NewRouter()
	router.Mount("/", routes.WebRoutes())
	router.Mount("/api/v1", routes.V1Api())

	srv := &http.Server{
		Addr:         cfg.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine so we can listen for shutdown signals
	go func() {
		log.Info("listening", "addr", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server error", "err", err)
		}
	}()

	// Block until we receive SIGINT or SIGTERM
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("shutting down gracefully…")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("forced shutdown", "err", err)
	}

	log.Info("server stopped")
}
