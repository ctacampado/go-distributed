package uauth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

//clientError makes it easier to return an error
func clientError(w *http.ResponseWriter, err *error, status int) {
	http.Error(*w, (*err).Error(), status)
}

// InitRouter initializes all the routes together with their respective
// handlers
func (a *UAuth) InitRouter() {
	r := a.Router
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// GET
	r.Get("/", a.index)
	// POST
	r.Post("/register", a.handleRegister)
	r.Post("/login", a.handleLogin)
	r.Post("/verify", a.handleVerify)
	r.Post("/disable", a.handleDisable)
	r.Post("/revoke", a.handleRevoke)
	r.Post("/delete", a.handleDelete)
}

// Index placeholder
func (a *UAuth) index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

// handleRegister accepts a body messge containing usernamer and password.
// The password is then salted and hashed the HashAndSalt function that is
// based on bcrypt (more on go-distributed/pkg/password).
func (a *UAuth) handleRegister(w http.ResponseWriter, r *http.Request) {

	c, sc, err := DecodeUserCredentials(r.Body)
	if err != nil {
		clientError(&w, &err, sc)
		return
	}

	sc, err = a.registerUserCredentials(c)
	// TODO: send verification link to email

	w.WriteHeader(sc)
	w.Write([]byte(`{"message": "Registration Success!"}`))
}

// HandleLogin accepts a body message containing username and password input
// from user in order to authenticate.
func (a *UAuth) handleLogin(w http.ResponseWriter, r *http.Request) {
	var id string
	c, sc, err := DecodeUserCredentials(r.Body)
	if err != nil {
		clientError(&w, &err, sc)
		return
	}

	id, sc, err = a.checkLoginCredentials(c.Username, c.Password)
	if err != nil {
		clientError(&w, &err, sc)
		return
	}

	w.WriteHeader(sc)
	w.Write([]byte(`{"id": "` + id + `","message": "Login Success!"}`))
}

// verify a user account
func (a *UAuth) handleVerify(w http.ResponseWriter, r *http.Request) {
	// check if verification code is valid
	// if valid, change user status to verified and return success
	// else return fail
}

// disable a user account
func (a *UAuth) handleDisable(w http.ResponseWriter, r *http.Request) {
	// check if authenticated and proper authorization
	// disable user
	// return success else return fail
}

// revoke a user account
func (a *UAuth) handleRevoke(w http.ResponseWriter, r *http.Request) {
	// check if authenticated and proper authorization
	// revoke user
	// return success else return fail
}

// delete a user account
func (a *UAuth) handleDelete(w http.ResponseWriter, r *http.Request) {
	// check if authenticated and proper authorization
	// delete user
	// return success else return fail
}
