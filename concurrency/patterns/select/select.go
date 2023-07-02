package main

import "fmt"

func main() {
	q := make(chan bool)
	select {
	case _ = <-q:
		fmt.Println("msg from queue")
	default:
		fmt.Println("default case")
	}
}
