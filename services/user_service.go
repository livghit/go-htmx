package services

import "net/http"

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mount this using r."))
}
