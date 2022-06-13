package entities

import (
	"math/rand"
	"time"
)

type Player struct {
	name string
	pos  int
}

// NewPlayer returns a new player
func NewPlayer(name string) *Player {
	return &Player{
		name: name,
		pos:  0,
	}
}

func (p *Player) getPosition() int {
	return p.pos
}

func (p *Player) setPosition(pos int) {
	p.pos = pos
}

// rollDice returns a random number between 1 - 6
func (p *Player) rollDice() int {
	rand.Seed(time.Now().UnixNano())
	return 1 + rand.Intn(6)
}
