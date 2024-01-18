package main

import (
	"fmt"
	// "golang-react-todo/router"
	"log"
	"net/http"

	"example.com/myserver/router"
)

func main() {
	r := router.Router()
	fmt.Println("starting the server on port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
