package main

import (
	"github.com/subosito/gotenv"

	auth "go-distributed/internal/auth"
)

func init() {
	gotenv.Load()
}

const port = "8080"

func main() {
	a := auth.NewAuth()
	a.StartAuthServer("auth server listening at localhost:", port)
}
