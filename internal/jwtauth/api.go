package jwtauth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

// InitRouter initializes all the routes together with their respective
// handlers
func (j *JWTServer) InitRouter() {
	r := j.Router
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", j.index)
	r.Post("/new", j.createNewToken)
	r.Post("/renew", j.renewToken)
	r.Post("/authenticate", j.authenticate)
	r.Delete("/delete", j.deleteToken)
}

// index placeholder
func (j *JWTServer) index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

// createNewToken placeholder
func (j *JWTServer) createNewToken(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "createNewToken!")
}

// renewToken placeholder
func (j *JWTServer) renewToken(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "createNewToken!")
}

// authenticate placeholder
func (j *JWTServer) authenticate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "authenticate!")
}

// deleteToken placeholder
func (j *JWTServer) deleteToken(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "deleteToken!")
}
