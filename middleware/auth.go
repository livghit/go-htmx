// Package middleware provides HTTP middleware for the application.
package middleware

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userIDKey contextKey = "userID"

// Protect is middleware that requires a valid JWT cookie.
// Unauthenticated requests are redirected to /login.
func Protect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, ok := validateCookie(r)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, setUserCtx(r, userID))
	})
}

// Optional loads the authenticated user into the request context if a valid
// JWT cookie is present, but never redirects. Use this on public pages that
// behave differently for logged-in users (e.g., show "Log out" in nav).
func Optional(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if userID, ok := validateCookie(r); ok {
			r = setUserCtx(r, userID)
		}
		next.ServeHTTP(w, r)
	})
}

// UserFromCtx retrieves the authenticated user's ID from the request context.
// Returns (0, false) if the user is not authenticated.
func UserFromCtx(r *http.Request) (int64, bool) {
	id, ok := r.Context().Value(userIDKey).(int64)
	return id, ok
}

// IsAuthenticated reports whether the request has a valid auth token.
func IsAuthenticated(r *http.Request) bool {
	_, ok := UserFromCtx(r)
	return ok
}

// SetCookie writes a signed JWT to an HttpOnly cookie.
// Call this after a successful login to establish the session.
func SetCookie(w http.ResponseWriter, userID int64) error {
	token, err := generateToken(userID)
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(jwtExpiry()),
	})
	return nil
}

// ClearCookie removes the auth cookie (logout).
func ClearCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "auth",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
}

// --- internal helpers ---

func validateCookie(r *http.Request) (int64, bool) {
	c, err := r.Cookie("auth")
	if err != nil {
		return 0, false
	}
	return validateToken(c.Value)
}

func setUserCtx(r *http.Request, userID int64) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), userIDKey, userID))
}

func jwtSecret() []byte {
	s := os.Getenv("JWT_SECRET")
	if s == "" {
		s = "dev-secret-change-in-production"
	}
	return []byte(s)
}

func jwtExpiry() time.Duration {
	return 24 * time.Hour
}

func generateToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"uid": userID,
		"exp": time.Now().Add(jwtExpiry()).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret())
}

func validateToken(tokenStr string) (int64, bool) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret(), nil
	})
	if err != nil || !token.Valid {
		return 0, false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, false
	}
	uid, ok := claims["uid"].(float64)
	if !ok {
		return 0, false
	}
	return int64(uid), true
}
