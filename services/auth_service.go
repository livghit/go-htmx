package services

import (
	"errors"
	"net/http"

	"github.com/livghit/go-htmx/db/models"
	"github.com/livghit/go-htmx/flash"
	"github.com/livghit/go-htmx/htmx"
	"github.com/livghit/go-htmx/middleware"
	"github.com/livghit/go-htmx/validation"
	"github.com/livghit/go-htmx/views/pages"
	"golang.org/x/crypto/bcrypt"
)

// ShowLogin renders the login page.
func ShowLogin(w http.ResponseWriter, r *http.Request) {
	flashes := flash.Get(w, r)
	Render(w, r, pages.LoginPage(flashes, validation.Errors{}))
}

// HandleLogin processes a login form submission.
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")

	errs := make(validation.Errors)
	validation.Required(username, "username", errs)
	validation.Required(password, "password", errs)

	if !errs.OK() {
		Render(w, r, pages.LoginPage(nil, errs))
		return
	}

	user, err := models.GetByUsername(username)
	if errors.Is(err, models.ErrNotFound) {
		errs["username"] = "invalid username or password"
		Render(w, r, pages.LoginPage(nil, errs))
		return
	}
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		errs["username"] = "invalid username or password"
		Render(w, r, pages.LoginPage(nil, errs))
		return
	}

	if err := middleware.SetCookie(w, user.ID); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	flash.Success(w, r, "Welcome back, "+user.Username+"!")

	// HTMX requests use HX-Redirect; regular forms use a normal 303 redirect.
	if htmx.IsHTMX(r) {
		htmx.Redirect(w, "/")
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// ShowRegister renders the registration page.
func ShowRegister(w http.ResponseWriter, r *http.Request) {
	flashes := flash.Get(w, r)
	Render(w, r, pages.RegisterPage(flashes, validation.Errors{}))
}

// HandleRegister processes a registration form submission.
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirm := r.FormValue("password_confirmation")

	errs := make(validation.Errors)
	validation.Required(username, "username", errs)
	validation.MinLen(username, "username", 3, errs)
	validation.Required(email, "email", errs)
	validation.IsEmail(email, "email", errs)
	validation.Required(password, "password", errs)
	validation.MinLen(password, "password", 8, errs)
	validation.Matches(password, confirm, "password_confirmation", errs)

	if !errs.OK() {
		Render(w, r, pages.RegisterPage(nil, errs))
		return
	}

	// Check uniqueness
	if _, err := models.GetByUsername(username); !errors.Is(err, models.ErrNotFound) {
		errs["username"] = "username is already taken"
		Render(w, r, pages.RegisterPage(nil, errs))
		return
	}
	if _, err := models.GetByEmail(email); !errors.Is(err, models.ErrNotFound) {
		errs["email"] = "email is already registered"
		Render(w, r, pages.RegisterPage(nil, errs))
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	user, err := models.Create(username, email, string(hashed))
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := middleware.SetCookie(w, user.ID); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	flash.Success(w, r, "Account created! Welcome, "+user.Username+"!")

	if htmx.IsHTMX(r) {
		htmx.Redirect(w, "/")
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// HandleLogout clears the auth cookie and redirects to the login page.
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	middleware.ClearCookie(w)
	flash.Info(w, r, "You have been logged out.")

	if htmx.IsHTMX(r) {
		htmx.Redirect(w, "/login")
		w.WriteHeader(http.StatusOK)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
