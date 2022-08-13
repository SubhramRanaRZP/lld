package entities

import (
	"fmt"
	"lld/elevator/src/enums"
	"lld/elevator/src/interfaces"
	"time"
)

type Elevator struct {
	// metadata
	id  string
	org Organisation

	// data
	cars        []*Car
	floors      []*Floor
	scheduler   interfaces.IScheduler
	alarm       *Alarm
	callService interfaces.ICallService
}

func (e *Elevator) GoDown(currentFloor *Floor) {
	car := e.scheduler.Schedule(&MoveRequest{
		sourceFloor: currentFloor,
		dir:         enums.RequestDirectionGoDown,
		time:        time.Now(),
	})
	car.AddStoppageFloor(currentFloor)
}

func (e *Elevator) GoUp(currentFloor *Floor) {
	car := e.scheduler.Schedule(&MoveRequest{
		sourceFloor: currentFloor,
		dir:         enums.RequestDirectionGoUp,
		time:        time.Now(),
	})
	car.AddStoppageFloor(currentFloor)
}

// CloseDoor will be triggered from the car's inner dashboard
func (e *Elevator) CloseDoor(c *Car) {
	if c.state == enums.CarStateTemporaryStop {
		for c.currentWt > c.maxWtCap {
			fmt.Println("over weight alarm...")
			time.Sleep(time.Duration(500) * time.Millisecond)
		}
		c.door.Close()
		c.currentFloorDoor.Close()
	}
}

// OpenDoor will be triggered from the car's inner dashboard
func (e *Elevator) OpenDoor(c *Car) {
	if c.state == enums.CarStateTemporaryStop {
		c.door.Open()
		c.currentFloorDoor.Open()
	}
}

func (e *Elevator) CallEmergency() {
	e.callService.Call()
}

// StopCar api will be consumed by admin only which is for maintenance purpose
func (e *Elevator) StopCar(c *Car) {
	c.state = enums.CarStateStopped
}

// StartCar api will be consumed by admin only which is called after maintenance
func (e *Elevator) StartCar(c *Car) {
	c.state = enums.CarStateTemporaryStop
}

func (e *Elevator) AddEmergencyPhoneNumber(num string) {
	e.callService.AddEmergencyPhoneNumber(num)
}

// AdminCallEmergency will be consumed by people
func (e *Elevator) AdminCallEmergency() {
	e.callService.Call()
}
