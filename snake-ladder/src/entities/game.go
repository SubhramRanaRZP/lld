package entities

import (
	"fmt"
	"lld/snake-ladder/ioHandler"
)

type IGame interface {
	Play()
}

type game struct {
	board     *board
	playerCnt int
	players   []*Player
}

// NewGame returns a new game
func NewGame(b *board, playerCnt int, players []*Player) IGame {
	return &game{
		board:     b,
		playerCnt: playerCnt,
		players:   players,
	}
}

func (g *game) Play() {
	for playerIndex := 0; ; playerIndex = (playerIndex + 1) % g.playerCnt {
		currentPlayer := g.players[playerIndex]
		currentPos := currentPlayer.getPosition()
		step := currentPlayer.rollDice()
		dest := g.board.move(currentPos, step)
		currentPlayer.setPosition(dest)

		ioHandler.Write(ioHandler.OutputFilePath,
			fmt.Sprintf("%v rolled a %v and moved from %v to %v",
				currentPlayer.name,
				step,
				currentPos,
				dest))

		if dest == g.board.destination {
			ioHandler.Write(ioHandler.OutputFilePath,
				fmt.Sprintf("%v wins the game", currentPlayer.name))
			break
		}
		// time.Sleep(100*time.Millisecond)
	}
}
