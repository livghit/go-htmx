package htmx

import "net/http"

// SetTrigger tells the client to trigger a custom event after the swap.
// The event is fired on the element that made the request.
//
//	htmx.SetTrigger(w, "userCreated")
func SetTrigger(w http.ResponseWriter, event string) {
	w.Header().Set("HX-Trigger", event)
}

// SetTriggerAfterSettle fires a custom event after the DOM has settled
// (after all CSS transitions have completed).
func SetTriggerAfterSettle(w http.ResponseWriter, event string) {
	w.Header().Set("HX-Trigger-After-Settle", event)
}

// SetTriggerAfterSwap fires a custom event after the swap step, but before
// the settle step.
func SetTriggerAfterSwap(w http.ResponseWriter, event string) {
	w.Header().Set("HX-Trigger-After-Swap", event)
}

// Redirect performs a client-side redirect without a full page reload.
// HTMX will follow the redirect, fetching the new URL and swapping its content.
// This is the HTMX equivalent of http.Redirect for HTMX requests.
//
//	htmx.Redirect(w, "/dashboard")
func Redirect(w http.ResponseWriter, url string) {
	w.Header().Set("HX-Redirect", url)
}

// Refresh triggers a full browser page refresh.
func Refresh(w http.ResponseWriter) {
	w.Header().Set("HX-Refresh", "true")
}

// PushURL pushes a new URL onto the browser history stack.
// Use "false" to prevent the URL from being pushed.
func PushURL(w http.ResponseWriter, url string) {
	w.Header().Set("HX-Push-Url", url)
}

// ReplaceURL replaces the current URL in the browser history without a new entry.
// Use "false" to prevent URL replacement.
func ReplaceURL(w http.ResponseWriter, url string) {
	w.Header().Set("HX-Replace-Url", url)
}

// Retarget overrides the target element for the response swap.
// The selector should be a CSS selector string (e.g., "#my-div").
func Retarget(w http.ResponseWriter, selector string) {
	w.Header().Set("HX-Retarget", selector)
}

// Reswap overrides the swap strategy for this response.
// Valid values: innerHTML, outerHTML, beforebegin, afterbegin, beforeend,
// afterend, delete, none.
func Reswap(w http.ResponseWriter, strategy string) {
	w.Header().Set("HX-Reswap", strategy)
}

// StopPolling responds with HTTP 286 to tell HTMX to stop polling.
// Use this in handlers targeted by hx-trigger="every Xs".
func StopPolling(w http.ResponseWriter) {
	w.WriteHeader(286)
}

// NoContent responds with 204 No Content, which tells HTMX to perform no swap.
// Useful for form submissions that should not update the DOM.
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
