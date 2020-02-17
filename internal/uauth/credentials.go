package uauth

import (
	"database/sql"
	"encoding/json"
	pw "go-distributed/pkg/password"
	"io"
	"net/http"

	"github.com/google/uuid"
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
	REVOKED                        //4
)

// UserCredentials structure is similar to the schema used in our DB
// to store user credentials.
type UserCredentials struct {
	ID             string     `json:"id,omitempty" db:"id"`               // unique ID for every user
	RoleID         int        `json:"roleid" db:"roleid"`                 // unique role id
	Username       string     `json:"username" db:"username"`             // username should be an email
	Password       string     `json:"password" db:"password"`             // salted and hashed
	Status         AcctStatus `json:"status,omitempty" db:"status"`       // says if an account is new, verivied, or disabled
	FailedAttempts int        `json:"fattempts,omitempty" db:"fattempts"` // number of failed sign-in attempts
}

// DecodeUserCredentials decodes the request body, represented by target
// into a variable of type UserCredentials
func DecodeUserCredentials(target io.ReadCloser) (uc *UserCredentials, sc int, err error) {
	uc = new(UserCredentials)
	err = json.NewDecoder(target).Decode(&uc)
	sc = http.StatusOK
	if err != nil {
		sc = http.StatusUnprocessableEntity
		return nil, sc, err
	}

	if uc.Username == "" || uc.Password == "" {
		sc = http.StatusBadRequest
		return nil, sc, err
	}
	return
}

func (a *UAuth) registerUserCredentials(c *UserCredentials) (int, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	c.ID = u.String()
	c.RoleID = a.AcctType
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

func (a *UAuth) checkLoginCredentials(un string, pwd string) (string, int, error) {
	var id sql.NullString
	var status AcctStatus
	var storedPW sql.NullString
	sqlStatement := `SELECT id,status,password FROM users WHERE username=($1)`
	err := a.DB.QueryRow(sqlStatement, un).Scan(&id, &status, &storedPW)
	switch {
	case err == sql.ErrNoRows || !storedPW.Valid:
		return id.String, http.StatusNotFound, err
	case err != nil:
		return id.String, http.StatusBadRequest, err
	case status == DISABLED:
		return id.String, http.StatusConflict, err
	case status == REVOKED:
		return id.String, http.StatusUnauthorized, err
	default:
		err = pw.ComparePasswords(storedPW.String, pwd)
		if err != nil {
			return id.String, http.StatusUnauthorized, err
		}
	}
	return id.String, http.StatusOK, nil
}
