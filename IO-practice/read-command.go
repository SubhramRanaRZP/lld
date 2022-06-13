// this file shows how to process commands in an LLD question
package main

import (
	"fmt"
)

type commandType int

const (
	checkout commandType = iota
	login
	logout
)

func (cmd commandType) toString() string {
	switch cmd {
	case checkout:
		return "checkout"
	case login:
		return "login"
	case logout:
		return "logout"
	default:
		return "unknown"
	}
}

func main() {
	commands := []string{"checkout", "logout", "login", "reserve", "checkout"}
	for _, cmd := range commands {
		do(cmd)
	}
}

func do(cmd string) {
	switch cmd {
	case checkout.toString():
		fmt.Println("checkout processing")
	case login.toString():
		fmt.Println("login processing")
	case logout.toString():
		fmt.Println("logout processing")
	default:
		fmt.Println("invalid command")
	}
}
