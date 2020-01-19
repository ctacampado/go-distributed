package auth

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	pw "go-distributed/pkg/password"

	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
)

//clientError makes it easier to return an error
func clientError(w *http.ResponseWriter, err *error, status int) {
	http.Error(*w, (*err).Error(), status)
}

// InitRouter initializes all the routes together with their respective
// handlers
func (a *Auth) InitRouter() {
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

	r.Get("/", a.index)
	r.Post("/register", a.handleRegister)
	r.Post("/login", a.handleLogin)
	r.Delete("/user", a.handleDelete)
}

// Index placeholder
func (a *Auth) index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

//
func (a *Auth) handleDelete(w http.ResponseWriter, r *http.Request) {
	sc := http.StatusOK
	w.WriteHeader(sc)
	w.Write([]byte(`{"message": "User Deleted!"}`))
}

// HandleRegister accepts a body messge containing usernamer and password.
// The password is then salted and hashed the HashAndSalt function that is
// based on bcrypt (more on go-distributed/pkg/password).
func (a *Auth) handleRegister(w http.ResponseWriter, r *http.Request) {

	c, sc, err := DecodeUserCredentials(r.Body)
	if err != nil {
		clientError(&w, &err, sc)
		return
	}

	sc, err = a.registerUserCredentials(c)

	w.WriteHeader(sc)
	w.Write([]byte(`{"message": "Registration Success!"}`))
}

// HandleLogin accepts a body message containing username and password input
// from user in order to authenticate.
func (a *Auth) handleLogin(w http.ResponseWriter, r *http.Request) {
	c, sc, err := DecodeUserCredentials(r.Body)
	if err != nil {
		clientError(&w, &err, sc)
		return
	}

	sc, err = a.checkLoginCredentials(c.Username, c.Password)
	if err != nil {
		clientError(&w, &err, sc)
		return
	}

	w.WriteHeader(sc)
	w.Write([]byte(`{"message": "Login Success!"}`))
}

func (a *Auth) registerUserCredentials(c *UserCredentials) (int, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	c.ID = u.String()
	c.Status = NEW
	c.FailedAttempts = 0
	c.Password = pw.HashAndSalt([]byte(c.Password))

	sqlStatement := `INSERT INTO users VALUES ($1, $2, $3, $4, $5)`
	_, err = a.DB.Query(sqlStatement, c.ID, c.Username, c.Password, c.Status, c.FailedAttempts)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (a *Auth) checkLoginCredentials(un string, pwd string) (int, error) {
	var storedPW sql.NullString
	sqlStatement := `SELECT password FROM users WHERE username=($1)`
	err := a.DB.QueryRow(sqlStatement, un).Scan(&storedPW)
	switch {
	case err == sql.ErrNoRows || !storedPW.Valid:
		return http.StatusNotFound, err
	case err != nil:
		return http.StatusBadRequest, err
	default:
		err = pw.ComparePasswords(storedPW.String, pwd)
		if err != nil {
			return http.StatusUnauthorized, err
		}
	}
	return http.StatusOK, nil
}
