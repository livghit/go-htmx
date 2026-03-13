package services

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/livghit/go-htmx/db/models"
	"github.com/livghit/go-htmx/flash"
	"github.com/livghit/go-htmx/middleware"
	"github.com/livghit/go-htmx/views/pages"
)

// Render writes a templ component to the response.
// The content type is set to text/html automatically.
func Render(w http.ResponseWriter, r *http.Request, component templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := component.Render(r.Context(), w); err != nil {
		http.Error(w, "render error", http.StatusInternalServerError)
	}
}

// HandleHomepage renders the homepage.
func HandleHomepage(w http.ResponseWriter, r *http.Request) {
	flashes := flash.Get(w, r)

	username := ""
	if userID, ok := middleware.UserFromCtx(r); ok {
		if user, err := models.GetByID(userID); err == nil {
			username = user.Username
		}
	}

	Render(w, r, pages.HomePage(flashes, username))
}
