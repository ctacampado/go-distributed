package auth

import (
	"database/sql"
	"go-distributed/pkg/dbclient"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// Auth structure holds pointers necessary for the service to work.
// These pointers are for the router, and the database
type Auth struct {
	Router *chi.Mux
	DB     *sql.DB
}

//NewAuth initializes a Login object with a new router
func NewAuth() *Auth {
	au := &Auth{
		DB:     dbclient.NewDB(),
		Router: chi.NewRouter(),
	}
	au.InitRouter()
	return au
}

// StartAuthServer starts an http listener hosted at localhost:port
func (a *Auth) StartAuthServer(msg string, port string) {
	log.Println(msg + " " + port)
	log.Fatal(http.ListenAndServe(":"+port, a.Router))
}
