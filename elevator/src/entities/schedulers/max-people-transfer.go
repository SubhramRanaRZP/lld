package schedulers

import (
	"fmt"
	"lld/elevator/src/entities"
)

type MaxPeopleTransferScheduler struct {
	name string
}

func (s *MaxPeopleTransferScheduler) Schedule(r *entities.MoveRequest) *entities.Car {
	fmt.Printf("request: %+v has been scheduled according to %v",
		*r, s.name)
	return nil
}
