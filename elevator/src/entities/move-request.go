package entities

import (
	"lld/elevator/src/enums"
	"time"
)

type MoveRequest struct {
	sourceFloor *Floor
	dir         enums.RequestDirection
	time        time.Time
}
