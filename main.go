package main

import (
	"fmt"
	"strings"

	configAnalyser "github.com/kimdepypere/golangAPPS/interpreter/cisco"
	webserver "github.com/kimdepypere/golangAPPS/webserver"
)

func main() {
	testInterpreter()
}

func testInterpreter() {
	//output := strings.Join(configAnalyser.FindAllInMap()[configAnalyser.Interfaces], "\n")
	output := strings.Join(configAnalyser.FindAll()[configAnalyser.Interfaces], "\n")
	fmt.Print(output)
}

func testWebserver() {
	webserver.TestRouter()
}
