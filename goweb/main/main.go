package main

import (
	"fmt"
	"log"
	"net/http"

	mux "github.com/kimdepypere/golangAPPS/go-router"
)

func main() {
	fmt.Println("Starting webserver")
	router := mux.NewRouter()
	router.GET("/welcome", getWelcome)
	router.GET("/test", getTest)
	router.GET("/", getDefault)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Page not found.")
}

func getWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to this website</h1>")
}

func getTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a testsite.")
}
