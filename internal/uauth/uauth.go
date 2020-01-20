package uauth

import (
	"database/sql"
	"go-distributed/pkg/dbclient"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

// UAuth structure holds pointers necessary for the User Authentication
// service to work. These pointers are for the router, and the database
type UAuth struct {
	Router *chi.Mux
	DB     *sql.DB
}

// NewUAuth initializes a User Auth Server
func NewUAuth() *UAuth {
	params := &dbclient.DBParams{
		Addr: dbclient.DBAddr{
			DBname: os.Getenv("LOGIN_DB_NAME"),
			Host:   os.Getenv("LOGIN_DB_HOST"),
			Port:   os.Getenv("LOGIN_DB_PORT"),
		},
		User:     os.Getenv("LOGIN_DB_USER"),
		Password: os.Getenv("LOGIN_DB_PASSWORD"),
		Sslmode:  os.Getenv("LOGIN_DB_SSL_MODE"),
	}
	au := &UAuth{
		DB:     dbclient.NewDB(params),
		Router: chi.NewRouter(),
	}
	au.InitRouter()
	return au
}

// StartUAuthServer starts an http listener hosted at localhost:port
func (a *UAuth) StartUAuthServer(msg string, port string) {
	log.Println(msg + " " + port)
	log.Fatal(http.ListenAndServe(":"+port, a.Router))
}
