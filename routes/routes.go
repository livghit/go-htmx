package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/livghit/go-htmx/middleware"
	"github.com/livghit/go-htmx/services"
)

// WebRoutes returns the main chi router with all web routes registered.
func WebRoutes() *chi.Mux {
	r := chi.NewRouter()

	// Standard middleware
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Recoverer)

	// Public routes
	r.With(middleware.Optional).Get("/", services.HandleHomepage)

	r.Get("/login", services.ShowLogin)
	r.Post("/login", services.HandleLogin)

	r.Get("/register", services.ShowRegister)
	r.Post("/register", services.HandleRegister)

	r.Post("/logout", services.HandleLogout)

	// Protected routes — require a valid auth cookie
	r.Group(func(r chi.Router) {
		r.Use(middleware.Protect)
		// r.Get("/dashboard", services.HandleDashboard)
		// Add more protected routes here
	})

	// Static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return r
}

// V1Api returns a chi.Mux for the /api/v1 route group.
func V1Api() *chi.Mux {
	r := chi.NewRouter()
	r.Use(chiMiddleware.RequestID)

	// API routes require auth
	r.Group(func(r chi.Router) {
		r.Use(middleware.Protect)
		r.Get("/users", services.GetAllUsers)
	})

	return r
}
