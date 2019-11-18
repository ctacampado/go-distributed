package controllers

import (
	"database/sql"
	"encoding/json"
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
	router.GET("/", Index)
	router.POST("/register", HandleRegister)

	return router
}

//Index -
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!")
}

//HandleRegister -
func HandleRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := Credentials{}

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(c)

	resp, _ := json.Marshal(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(resp)
}
