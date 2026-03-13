// Package htmx provides helpers for working with HTMX request headers and
// response headers in standard net/http handlers.
//
// HTMX sends special request headers to describe the state of the browser.
// Use the functions in this file to inspect those headers and branch your
// handler logic (e.g., return a partial instead of a full page).
package htmx

import "net/http"

// IsHTMX reports whether the request was initiated by HTMX
// (i.e., the HX-Request header is "true").
func IsHTMX(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}

// IsBoosted reports whether the request was triggered by hx-boost.
func IsBoosted(r *http.Request) bool {
	return r.Header.Get("HX-Boosted") == "true"
}

// IsHistoryRestore reports whether HTMX is restoring history
// from its local cache (the browser back button was used).
func IsHistoryRestore(r *http.Request) bool {
	return r.Header.Get("HX-History-Restore-Request") == "true"
}

// Target returns the id of the target element that HTMX will swap into.
// Empty string if not set.
func Target(r *http.Request) string {
	return r.Header.Get("HX-Target")
}

// TriggerName returns the name attribute of the element that triggered the
// request. Useful for differentiating multiple inputs in a form.
func TriggerName(r *http.Request) string {
	return r.Header.Get("HX-Trigger-Name")
}

// TriggerID returns the id of the element that triggered the request.
func TriggerID(r *http.Request) string {
	return r.Header.Get("HX-Trigger")
}

// CurrentURL returns the current URL of the browser when the request was made.
func CurrentURL(r *http.Request) string {
	return r.Header.Get("HX-Current-URL")
}

// Prompt returns the value entered by the user in an hx-prompt dialog.
func Prompt(r *http.Request) string {
	return r.Header.Get("HX-Prompt")
}
