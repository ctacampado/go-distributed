package controllers

import (
	"database/sql"
	"fmt"
	dbclient "go-distributed/pkg/dbclient"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

//Credentials -
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Login -
type Login struct {
	Router *httprouter.Router
	DB     *sql.DB
}

//NewLogin -
func NewLogin() *Login {
	return &Login{
		Router: initRouter(),
		DB: dbclient.InitDB(&dbclient.DBParams{
			Addr: dbclient.DBAddr{
				DBname: os.Getenv("DB_NAME"),
				Host:   os.Getenv("DB_HOST"),
				Port:   os.Getenv("DB_PORT"),
			},
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Sslmode:  os.Getenv("DB_SSL_MODE"),
		}),
	}
}

func initRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", index)

	return router
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
