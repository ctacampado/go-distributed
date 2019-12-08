package password

import (
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
// to the one that is stored in the DB. Since we'll be getting
// the hashed password from the DB it will be a string so we'll
// need to convert it to a byte slice
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
