package main

import (
	"fmt"
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
	fmt.Println(params)
	/*
		params := &dbclient.DBParams{
			Addr: dbclient.DBAddr{
				DBname: "logindb",
				Host:   "127.0.0.1",
				Port:   "5432",
			},
			User:     "postgres",
			Password: "postgres",
			Sslmode:  "disable",
		}*/
	login := controller.NewLogin(params)
	defer login.DB.Close()
	fmt.Printf("%+v\n", login.DB.Stats())
	//setup an http server
	//login.InitRoutes()
	//listen for server connections
	//log.Fatal(http.ListenAndServe(":3000", login.Router))
}
