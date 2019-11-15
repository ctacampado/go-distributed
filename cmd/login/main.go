package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/subosito/gotenv"

	controller "go-distributed/internal/controllers"
	"go-distributed/pkg/dbclient"
)

//Credentials -
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	gotenv.Load()
}

const port = "8080"

func main() {
	//initialize db

	params := &dbclient.DBParams{
		Addr: dbclient.DBAddr{
			DBname: os.Getenv("DB_NAME"),
			Host:   os.Getenv("DB_HOST"),
			Port:   os.Getenv("DB_PORT"),
		},
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Sslmode:  os.Getenv("DB_SSL_MODE"),
	}
	login := controller.NewLogin(params)
	defer login.DB.Close()

	login.InitRoutes()
	fmt.Println("login server listening at localhost:", port)
	log.Fatal(http.ListenAndServe(":"+port, login.Router))
}
