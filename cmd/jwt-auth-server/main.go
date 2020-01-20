package main

import (
	"go-distributed/internal/jwtauth"
	"os"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	port := os.Getenv("JWT_PORT")
	j := jwtauth.NewJWTServer()
	j.StartJWTServer("jwt server listening at localhost:", port)
}
