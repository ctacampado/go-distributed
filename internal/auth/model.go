package auth

import (
	"encoding/json"
	"io"
	"net/http"
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
	ID             string     `json:"id,omitempty" db:"id"`         // unique ID for every user
	Username       string     `json:"username" db:"username"`       // username should be an email
	Password       string     `json:"password" db:"password"`       // salted and hashed
	Status         AcctStatus `json:"Status,omitempty" db:"status"` // says if an account is new, verivied, or disabled
	FailedAttempts int        `db:"fattempts,omitempty"`            // number of failed sign-in attempts
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
