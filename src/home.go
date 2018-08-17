package main

import (
	"net/http"
)

const loginLink = `<html><body><a href="/login">Log in with Facebook</a></body></html>`

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(loginLink))
}
