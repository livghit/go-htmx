package services

import "net/http"

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Eywa you wish we had users XD."))
}
