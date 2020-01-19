package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/subosito/gotenv"

	login "go-distributed/internal/login"
	"go-distributed/pkg/dbclient"
)

func init() {
	gotenv.Load()
}

const port = "8080"

func main() {
	l := login.NewLogin()

	l.DB = dbclient.NewDB()
	defer l.DB.Close()

	l.InitRouter()
	fmt.Println("login server listening at localhost:", port)
	log.Fatal(http.ListenAndServe(":"+port, l.Router))
}
