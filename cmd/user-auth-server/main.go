package main

import (
	"github.com/subosito/gotenv"

	uauth "go-distributed/internal/uauth"
)

func init() {
	gotenv.Load()
}

const port = "8080"

func main() {
	a := uauth.NewUAuth()
	a.StartUAuthServer("auth server listening at localhost:", port)
}
