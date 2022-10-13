package middlewares

import (
	"net/http"
)

const (
	USERNAME = "admin"
	PASSWORD = "secret"
)

func Auth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()

	if !ok {
		w.Write([]byte(`something went wrong`))
		return false
	}

	isValid := (username == USERNAME) && (password == PASSWORD)

	if !isValid {
		w.Write([]byte(`wrong password or username`))
		return false
	}

	return true
}

func AllowOnlyGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte(`only get is allowed`))
		return false
	}	

	return true
}