package main

import (
	"fmt"

	"github.com/subosito/gotenv"

	controller "go-distributed/internal/controllers"
)

//Credentials -
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	gotenv.Load()
}

func main() {
	//initialize db
	login := controller.NewLogin()
	defer login.DB.Close()
	fmt.Printf("%+v\n", login.DB.Stats())
	//setup an http server

	//listen for server connections
}
