package tapps

import "fmt"

func hello(message string) string {
	if message == "" {
		fmt.Println("no message")
		return "Nothing to declare"
	}
	fmt.Println("a message will be given back")
	return "the message was: " + message
}

func goodbye() string {
	return "goodbye"
}
