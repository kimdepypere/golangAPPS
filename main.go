package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	mux "github.com/kimdepypere/golangAPPS/go-router"
	configAnalyser "github.com/kimdepypere/golangAPPS/interpreter/cisco"
)

func main() {
	testInterpreter()
}

func testInterpreter() {
	output := strings.Join(configAnalyser.FindInterfaces()[:], "\n")
	fmt.Print(output)
}

func testRouter() {
	fmt.Println("Starting webserver")
	router := mux.NewRouter()
	router.GET("/", getDynamic)
	router.GET("", getDynamic)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getDynamic(w http.ResponseWriter, r *http.Request) {

	shortpath := strings.TrimPrefix(r.URL.Path, "/")

	switch shortpath {
	case "welcome":
		fmt.Fprintf(w, "Welcome to my website part II")
	case "test":
		fmt.Fprintf(w, "This is a second testpage")
	default:
		fmt.Fprintf(w, "this is a second default page")
	}
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
