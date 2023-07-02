/*
This is a method known as "communicating by sharing" --> may cause deadlock
when all the philosopher will acquire a lock on the left fork

Solution --> "communicate by sharing". keep a coordinate goroutine
that will synchronise everyone's request to access a lock on necessary mutex
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numPhilosopher = 5
	eatTime        = 1000 // in milliseconds
)

var (
	forks []sync.Mutex
)

func philosopherEating(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	forks[id].Lock()
	forks[(id+1)%numPhilosopher].Lock()
	defer forks[id].Unlock()
	defer forks[(id+1)%numPhilosopher].Unlock()

	fmt.Printf("phil %v started eating\n", id)
	time.Sleep(eatTime * time.Millisecond) // to simulate eating process
	fmt.Printf("phil %v finished eating\n", id)
}

func main() {
	forks = make([]sync.Mutex, numPhilosopher)

	wg := &sync.WaitGroup{}
	wg.Add(numPhilosopher)
	for i := 0; i < numPhilosopher; i++ {
		go philosopherEating(i, wg)
	}

	wg.Wait()

	fmt.Println("all philosopher finished eating !")
}
