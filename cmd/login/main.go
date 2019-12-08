package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/subosito/gotenv"

	controller "go-distributed/internal/controllers"
	"go-distributed/pkg/dbclient"
)

func init() {
	gotenv.Load()
}

const port = "8080"

func main() {
	login := controller.NewLogin()

	login.DB = dbclient.NewDB()
	defer login.DB.Close()

	login.InitRouter()
	fmt.Println("login server listening at localhost:", port)
	log.Fatal(http.ListenAndServe(":"+port, login.Router))
}
