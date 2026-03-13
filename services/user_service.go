package services

import (
	"net/http"

	"github.com/livghit/go-htmx/db/models"
	"github.com/livghit/go-htmx/middleware"
)

// GetAllUsers returns all users as JSON. Requires authentication.
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.UserFromCtx(r)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// TODO: add pagination, filtering
	_ = models.User{} // placeholder until full implementation
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"users":[]}`))
}
