package uauth

import (
	"database/sql"
	"go-distributed/pkg/dbclient"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
)

// UAuth structure holds pointers necessary for the User Authentication
// service to work. These pointers are for the router, and the database
type UAuth struct {
	Router   *chi.Mux
	DB       *sql.DB
	AcctType int // 1-User 2-Admin
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

	accttype, err := strconv.Atoi(os.Getenv("LOGIN_TYPE"))
	if err != nil {
		log.Fatal("failed to get LOGIN_TYPE")
	}

	ua := &UAuth{
		DB:       dbclient.NewDB(params),
		Router:   chi.NewRouter(),
		AcctType: accttype,
	}
	ua.InitRouter()
	return ua
}

// StartUAuthServer starts an http listener hosted at localhost:port
func (a *UAuth) StartUAuthServer(msg string, port string) {
	log.Println(msg + " " + port)
	log.Fatal(http.ListenAndServe(":"+port, a.Router))
}
