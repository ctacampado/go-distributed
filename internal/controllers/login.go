package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	uuid "github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	pw "go-distributed/pkg/password"
)

// AcctStatus is used as a custom type to
// easily differentiate between user
// account status
type AcctStatus int

// different account status
const (
	NEW      AcctStatus = iota + 1 //1
	VERIFIED                       //2
	DISABLED                       //3
)

// UserCredentials structure is similar to the schema used in our DB
// to store user credentials.
type UserCredentials struct {
	ID             string     `json:"id,omitempty" db:"id"`         //unique ID for every user
	Username       string     `json:"username" db:"username"`       //username should be an email
	Password       string     `json:"password" db:"password"`       //salted and hashed
	Status         AcctStatus `json:"Status,omitempty" db:"status"` // says if an account is new, verivied, or disabled
	FailedAttempts int        `db:"fattempts,omitempty"`            //number of failed sign-in attempts
}

// Login structure holds pointers necessary for the service to work.
// These pointers are for the router, and the database
type Login struct {
	Router *httprouter.Router
	DB     *sql.DB
}

//NewLogin initializes a Login object with a new router
func NewLogin() *Login {
	return &Login{
		Router: httprouter.New(),
	}
}

//clientError makes it easier to return an error
func clientError(w *http.ResponseWriter, err *error, status int) {
	http.Error(*w, (*err).Error(), status)
}

// InitRouter initializes all the routes together with their respective
// handlers
func (l *Login) InitRouter() {
	l.Router.GET("/", l.Index)
	l.Router.POST("/register", l.HandleRegister)
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
func (l *Login) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!")
}

// HandleRegister accepts a body messge containing usernamer and password.
// The password is then salted and hashed the HashAndSalt function that is
// based on bcrypt (more on go-distributed/pkg/password).
func (l *Login) HandleRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//store username password to database
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
