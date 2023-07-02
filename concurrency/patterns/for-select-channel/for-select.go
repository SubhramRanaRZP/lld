package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func process(pid string, q chan<- string) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		q <- fmt.Sprintf("%v : %v", pid, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func isAllProcessFinished(done chan bool) {
	wg.Wait()
	done <- true
}

func main() {
	wg = sync.WaitGroup{}
	q1 := make(chan string)
	q2 := make(chan string)
	done := make(chan bool)

	wg.Add(2)
	go isAllProcessFinished(done)
	go process("pid 1", q1)
	go process("pid 2", q2)

	allDone := false
	for !allDone {
		select {
		case msg1 := <-q1:
			fmt.Println(msg1)
		case msg2 := <-q2:
			fmt.Println(msg2)
		case _ = <-done:
			fmt.Println("all processing done")
			allDone = true
		}
	}

	fmt.Println("some post processing")
	time.Sleep(time.Second)
}
