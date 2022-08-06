package main

import (
	"lld/snake-ladder/ioHandler"
	entities2 "lld/snake-ladder/src/entities"
)

func main() {
	var snakeCnt, ladderCnt, playerCnt int
	var snakes, ladders [][]int
	var playerNames []string
	ioHandler.TakeInput(&snakeCnt, &ladderCnt, &snakes, &ladders, &playerCnt, &playerNames)

	var players []*entities2.Player
	for _, name := range playerNames {
		players = append(players, entities2.NewPlayer(name))
	}

	IGame := entities2.NewGame(
		entities2.NewBoard(100, snakeCnt, snakes, ladderCnt, ladders),
		playerCnt,
		players)

	IGame.Play()
}
