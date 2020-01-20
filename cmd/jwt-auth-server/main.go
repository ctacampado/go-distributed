package main

import (
	"go-distributed/internal/jwtauth"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

const port = "8081"

func main() {

	j := jwtauth.NewJWTServer()
	j.StartJWTServer("jwt server listening at localhost:", port)
}
