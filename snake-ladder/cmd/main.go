package main

import (
	"lld/snake-ladder/entities"
	"lld/snake-ladder/ioHandler"
)

func main() {
	var snakeCnt, ladderCnt, playerCnt int
	var snakes, ladders [][]int
	var playerNames []string
	ioHandler.TakeInput(&snakeCnt, &ladderCnt, &snakes, &ladders, &playerCnt, &playerNames)

	var players []*entities.Player
	for _, name := range playerNames{
		players = append(players, entities.NewPlayer(name))
	}

	IGame := entities.NewGame(
		entities.NewBoard(100, snakeCnt, snakes, ladderCnt, ladders),
		playerCnt,
		players)

	IGame.Play()
}
