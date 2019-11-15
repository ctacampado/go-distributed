package controllers

import (
	"database/sql"
	"fmt"
	dbclient "go-distributed/pkg/dbclient"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Login -
type Login struct {
	Router *httprouter.Router
	DB     *sql.DB
}

//NewLogin -
func NewLogin(params *dbclient.DBParams) *Login {

	return &Login{
		Router: httprouter.New(),
		DB:     dbclient.InitDB(params),
	}
}

//InitRoutes -
func (l *Login) InitRoutes() {
	l.Router.GET("/", Index)
}

//Index -
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
