package password

import (
	"errors"
	"log"

	bcrypt "golang.org/x/crypto/bcrypt"
)

// HashAndSalt uses GenerateFromPassword to hash & salt pwd
// MinCost is just an integer constant provided by the bcrypt
// package along with DefaultCost & MaxCost.
// The cost can be any value you want provided it isn't lower
// than the MinCost (4)
func HashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

// ComparePasswords compares the password input by the user
// to the one that is stored in the DB. both passwords
// has to be converted from string to bytes in order to use
// bcrypt's CompareHashAndPassword function
func ComparePasswords(hashedPwd string, plainPwd string) error {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		log.Println(err)
		return errors.New("invalid credentials")
	}

	return nil
}
