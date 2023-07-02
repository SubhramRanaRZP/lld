package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func dummy(counter int) {
	defer wg.Done()
	fmt.Println("I am dummy function for: ", counter)
}

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go dummy(i)
	}

	wg.Wait()
	fmt.Println("all process finished")
}
