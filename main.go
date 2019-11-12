package main

import (
	"fmt"
	"strings"

	configAnalyser "github.com/kimdepypere/golangAPPS/interpreter/cisco"
	webserver "github.com/kimdepypere/golangAPPS/webserver"
)

func main() {
	fmt.Print(testInterpreter())
}

func testInterpreter() string {
	output := strings.Join(configAnalyser.FindAll()[configAnalyser.Interfaces], "\n")
	return output
}

func testWebserver() {
	webserver.TestRouter()
}
