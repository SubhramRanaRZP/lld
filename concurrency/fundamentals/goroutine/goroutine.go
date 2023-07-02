package main

import (
	"fmt"
	"time"
)

func printName(name string) {
	for i := 0; i < 3; i++ {
		fmt.Println(name)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go printName("Subhram")
	go printName("Rana")

	time.Sleep(3 * time.Second)
	fmt.Println("done")
}
