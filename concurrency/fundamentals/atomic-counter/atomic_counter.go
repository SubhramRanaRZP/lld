package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var counter int64 = 0

func incrementBy1() {
	for j := 0; j < 100; j++ {
		//counter++ // not thread safe, ==> data race
		atomic.AddInt64(&counter, 1) // thread safe, ==> no data race
	}
}

func main() {
	fmt.Println("initial counter: ", counter)

	for i := 0; i < 50; i++ {
		go incrementBy1()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("final counter: ", counter)
}
