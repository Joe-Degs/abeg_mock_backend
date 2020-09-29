// Package controllers define functions that control how a route handles a request.
package controllers

import "net/http"

// Login controls the login route of server
func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login handler\n"))
}

// Signup controls the signup route of server
func Signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("signup handler\n"))
}
