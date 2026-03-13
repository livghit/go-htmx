// Package validation provides lightweight form and request validation helpers.
// It is intentionally simple — no reflection, no struct tags, just explicit checks.
//
// Usage:
//
//	errs := make(validation.Errors)
//	validation.Required(r.FormValue("email"), "email", errs)
//	validation.IsEmail(r.FormValue("email"), "email", errs)
//	validation.MinLen(r.FormValue("password"), "password", 8, errs)
//	if !errs.OK() {
//	    // render form again with errs
//	}
package validation

import (
	"fmt"
	"net/mail"
	"unicode/utf8"
)

// Errors maps field names to their first validation error message.
type Errors map[string]string

// OK reports whether there are no validation errors.
func (e Errors) OK() bool {
	return len(e) == 0
}

// First returns the first error message for a field, or empty string.
func (e Errors) First(field string) string {
	return e[field]
}

// Has reports whether there is an error for the given field.
func (e Errors) Has(field string) bool {
	_, ok := e[field]
	return ok
}

// Required checks that val is not empty. Adds an error under field if it is.
func Required(val, field string, errs Errors) {
	if errs.Has(field) {
		return
	}
	if len(val) == 0 {
		errs[field] = fmt.Sprintf("%s is required", field)
	}
}

// MinLen checks that val has at least min UTF-8 characters.
func MinLen(val, field string, min int, errs Errors) {
	if errs.Has(field) {
		return
	}
	if utf8.RuneCountInString(val) < min {
		errs[field] = fmt.Sprintf("%s must be at least %d characters", field, min)
	}
}

// MaxLen checks that val has at most max UTF-8 characters.
func MaxLen(val, field string, max int, errs Errors) {
	if errs.Has(field) {
		return
	}
	if utf8.RuneCountInString(val) > max {
		errs[field] = fmt.Sprintf("%s must be at most %d characters", field, max)
	}
}

// IsEmail checks that val is a syntactically valid email address.
func IsEmail(val, field string, errs Errors) {
	if errs.Has(field) {
		return
	}
	if _, err := mail.ParseAddress(val); err != nil {
		errs[field] = fmt.Sprintf("%s must be a valid email address", field)
	}
}

// Matches checks that a equals b (e.g., password == password_confirmation).
func Matches(a, b, field string, errs Errors) {
	if errs.Has(field) {
		return
	}
	if a != b {
		errs[field] = fmt.Sprintf("%s does not match", field)
	}
}

// Unique can be called with a boolean result from a DB existence check.
// If exists is true, it adds a "already taken" error.
func Unique(exists bool, field string, errs Errors) {
	if errs.Has(field) {
		return
	}
	if exists {
		errs[field] = fmt.Sprintf("%s is already taken", field)
	}
}
