package main

import (
	"fmt"
	"time"
)

func getNextNumber(q chan<- string) {
	for i := 0; i < 20; i++ {
		q <- fmt.Sprintf("number: %v", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	timer := time.NewTimer(1 * time.Second)

	q := make(chan string)
	go getNextNumber(q)

	timeOver := false
	for !timeOver {
		select {
		case msg := <-q:
			fmt.Println(msg)
		case <-timer.C:
			timeOver = true
		}
	}

	fmt.Println("time over")
}
