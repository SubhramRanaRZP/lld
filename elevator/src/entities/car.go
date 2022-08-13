package entities

import (
	"lld/elevator/src/enums"
	"lld/elevator/src/interfaces"
)

type Car struct {
	id                  string
	door                *Door
	state               enums.CarState
	maxWtCap, currentWt float64
	display             *Display
	stoppageFloors      []*Floor
	currentFloorDoor    *Door // this will be nil in the case where cat is in moving state
	callService         interfaces.ICallService
}

func (c *Car) OpenDoor() {
	c.door.Open()
}

func (c *Car) CloseDoor() {
	c.door.Close()
}

func (c *Car) AddStoppageFloor(f *Floor) {
	c.stoppageFloors = append(c.stoppageFloors, f)
}

// CallEmergency will be consumed by people
func (c *Car) CallEmergency() {
	c.callService.Call()
}
