package schedulers

import (
	"fmt"
	"lld/elevator/src/entities"
)

type CostOptimizerScheduler struct {
	name string
}

func (s *CostOptimizerScheduler) Schedule(r *entities.MoveRequest) *entities.Car {
	fmt.Printf("request: %+v has been scheduled according to %v",
		*r, s.name)
	return nil
}
