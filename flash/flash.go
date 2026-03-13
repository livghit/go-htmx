// Package flash provides HTTP cookie-based flash messages.
//
// Flash messages are short-lived notifications that survive exactly one
// redirect. They are stored in a signed cookie so no server-side session
// store is required.
//
// Usage in a handler:
//
//	flash.Success(w, r, "Welcome back!")
//	http.Redirect(w, r, "/", http.StatusSeeOther)
//
// Then in the destination handler, retrieve and clear them:
//
//	msgs := flash.Get(w, r)  // clears cookie automatically
package flash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const cookieName = "flash"

// Type represents the visual style of a flash message.
type Type string

const (
	TypeSuccess Type = "success"
	TypeError   Type = "error"
	TypeWarning Type = "warning"
	TypeInfo    Type = "info"
)

// Flash is a single flash message.
type Flash struct {
	Type    Type   `json:"t"`
	Message string `json:"m"`
}

// Success adds a success-level flash message.
func Success(w http.ResponseWriter, r *http.Request, msg string) {
	Set(w, r, TypeSuccess, msg)
}

// Error adds an error-level flash message.
func Error(w http.ResponseWriter, r *http.Request, msg string) {
	Set(w, r, TypeError, msg)
}

// Warning adds a warning-level flash message.
func Warning(w http.ResponseWriter, r *http.Request, msg string) {
	Set(w, r, TypeWarning, msg)
}

// Info adds an info-level flash message.
func Info(w http.ResponseWriter, r *http.Request, msg string) {
	Set(w, r, TypeInfo, msg)
}

// Set stores a flash message in a signed cookie.
// Multiple calls append to the existing flash cookie.
func Set(w http.ResponseWriter, r *http.Request, t Type, msg string) {
	existing := readRaw(r)
	existing = append(existing, Flash{Type: t, Message: msg})

	raw, err := marshal(existing)
	if err != nil {
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    raw,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}

// Get returns all pending flash messages and immediately clears the cookie.
// Call this once per request, typically in your layout or middleware.
func Get(w http.ResponseWriter, r *http.Request) []Flash {
	msgs := readRaw(r)
	if len(msgs) == 0 {
		return nil
	}
	// clear the cookie
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
	return msgs
}

// readRaw reads and verifies the flash cookie without clearing it.
func readRaw(r *http.Request) []Flash {
	c, err := r.Cookie(cookieName)
	if err != nil {
		return nil
	}
	msgs, err := unmarshal(c.Value)
	if err != nil {
		return nil
	}
	return msgs
}

func secret() []byte {
	s := os.Getenv("FLASH_SECRET")
	if s == "" {
		s = "dev-flash-secret"
	}
	return []byte(s)
}

func marshal(msgs []Flash) (string, error) {
	data, err := json.Marshal(msgs)
	if err != nil {
		return "", err
	}
	encoded := base64.URLEncoding.EncodeToString(data)
	sig := sign(encoded)
	return fmt.Sprintf("%s.%s", encoded, sig), nil
}

func unmarshal(raw string) ([]Flash, error) {
	parts := strings.SplitN(raw, ".", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid flash cookie format")
	}
	encoded, sig := parts[0], parts[1]
	if !hmac.Equal([]byte(sig), []byte(sign(encoded))) {
		return nil, fmt.Errorf("flash cookie signature mismatch")
	}
	data, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	var msgs []Flash
	if err := json.Unmarshal(data, &msgs); err != nil {
		return nil, err
	}
	return msgs, nil
}

func sign(value string) string {
	mac := hmac.New(sha256.New, secret())
	mac.Write([]byte(value))
	return base64.URLEncoding.EncodeToString(mac.Sum(nil))
}
