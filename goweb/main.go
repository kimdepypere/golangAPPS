package main

import (
	"fmt"
	"log"
	"net/http"

	mux "github.com/kimdepypere/golangAPPS/go-router"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to this website</h1>")
	log.Println("Accessed welcome page.")
}

func main() {
	fmt.Println("Starting webserver")
	router := mux.NewRouter()
	router.GET("/welcome", welcome)
	log.Fatal(http.ListenAndServe(":8080", router))
}
