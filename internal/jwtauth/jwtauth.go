package jwtauth

import (
	"database/sql"
	"go-distributed/pkg/dbclient"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

// JWTServer is a struct that defines a JWT Authentication Server
type JWTServer struct {
	Router *chi.Mux
	DB     *sql.DB
}

// NewJWTServer initializes a JWT Auth Server
func NewJWTServer() *JWTServer {
	params := &dbclient.DBParams{
		Addr: dbclient.DBAddr{
			DBname: os.Getenv("JWT_DB_NAME"),
			Host:   os.Getenv("JWT_DB_HOST"),
			Port:   os.Getenv("JWT_DB_PORT"),
		},
		User:     os.Getenv("JWT_DB_USER"),
		Password: os.Getenv("JWT_DB_PASSWORD"),
		Sslmode:  os.Getenv("JWT_DB_SSL_MODE"),
	}
	jas := &JWTServer{
		DB:     dbclient.NewDB(params),
		Router: chi.NewRouter(),
	}
	jas.InitRouter()
	return jas
}

// StartJWTServer starts an http listener hosted at localhost:port
func (j *JWTServer) StartJWTServer(msg string, port string) {
	log.Println(msg + " " + port)
	log.Fatal(http.ListenAndServe(":"+port, j.Router))
}
