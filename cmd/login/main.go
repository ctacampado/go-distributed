package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/subosito/gotenv"

	controller "go-distributed/internal/controllers"
)

func init() {
	gotenv.Load()
}

const port = "8080"

func main() {
	login := controller.NewLogin()
	defer login.DB.Close()

	fmt.Println("login server listening at localhost:", port)
	log.Fatal(http.ListenAndServe(":"+port, login.Router))
}
