/*
stateful goroutine pattern will be used here
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
	totPhilosopher = 5
	eatTime        = 1000 // milliseconds
)

type request struct {
	requiredForks []int
	respond       chan bool
}

// ---------------------coordinator-----------------------------
type coordinator struct {
	forks  []*sync.Mutex
	msgQue chan *request
}

func (c *coordinator) HandleRequests() {
	for {
		if req, more := <-c.msgQue; more {
			c.processRequest(req)
		} else {
			break
		}
	}
}
func (c *coordinator) getFork(id int) *sync.Mutex {
	return c.forks[id]
}
func (c *coordinator) processRequest(req *request) {
	for i, forkId := range req.requiredForks {
		if !c.getFork(forkId).TryLock() {
			// unlock all previous locks
			for j := 0; j <= i; j++ {
				c.getFork(req.requiredForks[j]).Unlock()
			}

			req.respond <- false
			return
		}
	}
	req.respond <- true
}

// ------------------------coordinator--------------------------

// ----------------- philosopher----------------
type philosopher struct {
	wg     *sync.WaitGroup
	id     int
	msgQue chan *request
}

func (p *philosopher) eat() {
	fmt.Printf("philosopher %v started eating\n", p.id)
	time.Sleep(eatTime * time.Millisecond)
	fmt.Printf("philosopher %v finished eating\n", p.id)
}
func (p *philosopher) manageEatingProcess() {
	defer p.wg.Done()

	req := &request{
		requiredForks: []int{p.id, (p.id + 1) % numPhilosopher},
		respond:       make(chan bool),
	}

	for {
		p.msgQue <- req
		resp := <-req.respond
		if resp {
			break
		}
	}
}

// ----------------- philosopher----------------

func main() {
	wg := &sync.WaitGroup{}
	msgQue := make(chan *request, totPhilosopher)

	// 1. initializing forks
	allForks := make([]*sync.Mutex, 0)
	for i := 0; i < numPhilosopher; i++ {
		allForks = append(allForks, &sync.Mutex{})
	}
	// 2. initializing msg queue
	msgQue := make(chan *request)

	coordinator := coordinator{
		forks:  allForks,
		msgQue: msgQue,
	}
}
