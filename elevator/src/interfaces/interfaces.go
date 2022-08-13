package interfaces

import "lld/elevator/src/entities"

// IScheduler facilitate Admin to select scheduler of his/her own choice which will schedule
// a car for a goUp/goDown request
type IScheduler interface {
	Schedule(r *entities.MoveRequest) *entities.Car
}

// ICallService can be used in future to integrate with some other calling service
type ICallService interface {
	Call()
	AddEmergencyPhoneNumber(num string)
}
