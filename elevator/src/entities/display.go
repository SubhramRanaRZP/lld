package entities

import (
	"fmt"
	"lld/elevator/src/enums"
)

type Display struct {
	currentFloor *Floor
	direction    enums.CarState
}

func (d *Display) DisplayCarData() {
	fmt.Printf("currentFloor: %v , car direction: %v",
		d.currentFloor.floorNumber, d.direction)
}
