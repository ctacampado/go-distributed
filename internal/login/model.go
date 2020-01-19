package login

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
