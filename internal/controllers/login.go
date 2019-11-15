package controllers

import (
	"database/sql"
	dbclient "go-distributed/pkg/dbclient"
	"os"
)

//Login -
type Login struct {
	DB *sql.DB
}

//NewLogin -
func NewLogin() *Login {
	params := dbclient.DBParams{
		Addr: dbclient.DBAddr{
			DBname: os.Getenv("DB_NAME"),
			Host:   os.Getenv("DB_HOST"),
			Port:   os.Getenv("DB_PORT"),
		},
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Sslmode:  os.Getenv("DB_SSL_MODE"),
	}
	return &Login{DB: dbclient.InitDB(params)}
}
