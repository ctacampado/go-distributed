package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	pw "go-distributed/pkg/password"

	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
)

// InitRouter initializes all the routes together with their respective
// handlers
func (l *Login) InitRouter() {
	r := l.Router
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", l.index)
	r.Post("/register", l.handleRegister)
	r.Post("/login", l.handleLogin)
}

//clientError makes it easier to return an error
func clientError(w *http.ResponseWriter, err *error, status int) {
	http.Error(*w, (*err).Error(), status)
}

func (l *Login) insertRegistrationToDB(c *UserCredentials) error {
	sqlStatement := `INSERT INTO users VALUES ($1, $2, $3, $4, $5)`
	_, err := l.DB.Query(sqlStatement, c.ID, c.Username, c.Password, c.Status, c.FailedAttempts)
	if err != nil {
		return err
	}

	return nil
}

//Index placeholder
func (l *Login) index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

// HandleRegister accepts a body messge containing usernamer and password.
// The password is then salted and hashed the HashAndSalt function that is
// based on bcrypt (more on go-distributed/pkg/password).
func (l *Login) handleRegister(w http.ResponseWriter, r *http.Request) {
	createReq := new(UserCredentials)
	err := json.NewDecoder(r.Body).Decode(&createReq)
	if err != nil {
		clientError(&w, &err, http.StatusUnprocessableEntity)
		return
	}

	if createReq.Username == "" || createReq.Password == "" {
		clientError(&w, &err, http.StatusBadRequest)
		return
	}

	u, err := uuid.NewRandom()
	if err != nil {
		clientError(&w, &err, http.StatusInternalServerError)
		return
	}
	createReq.ID = u.String()
	createReq.Status = NEW
	createReq.FailedAttempts = 0
	createReq.Password = pw.HashAndSalt([]byte(createReq.Password))

	err = l.insertRegistrationToDB(createReq)
	if err != nil {
		clientError(&w, &err, http.StatusInternalServerError)
		return
	}

	w.Write([]byte(`{"message": "Registration Success!"}`))
}

// HandleLogin accepts a body message containing username and password input
// from user in order to authenticate.
func (l *Login) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"message": "Login Success!"}`))
}
