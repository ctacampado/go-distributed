package main

import (
	uauth "go-distributed/internal/uauth"
	"os"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	port := os.Getenv("LOGIN_PORT")
	a := uauth.NewUAuth()
	a.StartUAuthServer("auth server listening at localhost:", port)
}
