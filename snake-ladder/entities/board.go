package entities

type board struct {
	destination         int
	snakeCnt, ladderCnt int
	snakes, ladders     map[int]int
}

// NewBoard returns a new board
func NewBoard(dest int, snakeCnt int, snakes [][]int,
	ladderCnt int, ladders [][]int) *board {
	b := board{
		destination: dest,
		snakeCnt:    snakeCnt,
		ladderCnt:   ladderCnt,
	}

	b.snakes = make(map[int]int)
	for _, snake := range snakes {
		b.snakes[snake[0]] = snake[1]
	}

	b.ladders = make(map[int]int)
	for _, ladder := range ladders {
		b.ladders[ladder[0]] = ladder[1]
	}

	return &b
}

// move returns the final destination after moving 'step' distance
// from the 'start' position considering the ladders and snakes
func (b *board) move(start int, step int) int {
	next := start + step
	if next == b.destination {
		return b.destination
	} else if next > b.destination {
		return start
	}

	return b.reachViaSnakeOrLadder(next)
}

// reachViaSnakeOrLadder gives the final place where the Player
// will reach only via snakes and ladders without rolling dice
func (b *board) reachViaSnakeOrLadder(start int) int {
	ladderDest, ladderAvailable := b.ladders[start]
	snakeDest, snakeAvailable := b.snakes[start]

	if !ladderAvailable && !snakeAvailable {
		return start
	}

	if ladderAvailable {
		return b.reachViaSnakeOrLadder(ladderDest)
	}

	return b.reachViaSnakeOrLadder(snakeDest)
}
