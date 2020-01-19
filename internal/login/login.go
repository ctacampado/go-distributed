package login

import (
	"database/sql"

	"github.com/go-chi/chi"
)

// Login structure holds pointers necessary for the service to work.
// These pointers are for the router, and the database
type Login struct {
	Router *chi.Mux
	DB     *sql.DB
}

//NewLogin initializes a Login object with a new router
func NewLogin() *Login {
	return &Login{
		Router: chi.NewRouter(),
	}
}
