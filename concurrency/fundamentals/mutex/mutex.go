package main

import (
	"fmt"
	"sync"
	"time"
)

type player struct {
	name          string
	score         int
	lastUpdatedAt time.Time

	mu sync.Mutex
}

func (p *player) print() {
	fmt.Println(map[string]interface{}{
		"name":       p.name,
		"score":      p.score,
		"updated_at": p.lastUpdatedAt,
	})
}

func main() {
	player1 := player{
		name:          "Subhram",
		score:         0,
		lastUpdatedAt: time.Now(),
		mu:            sync.Mutex{},
	}

	for i := 1; i <= 2000; i++ {
		go func(increment int) {
			// it would be a data race scenario in the absence of below 2 lines
			player1.mu.Lock()
			defer player1.mu.Unlock()

			player1.score = player1.score + increment
			player1.lastUpdatedAt = time.Now()
		}(i)
	}

	time.Sleep(3 * time.Second)
	player1.print() // score should be = (2000 * 2001)/2 = 2001000
}
