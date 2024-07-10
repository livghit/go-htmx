package services

import "net/http"

func HandleHomepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Homepage buddy"))
}
