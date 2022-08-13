package schedulers

import (
	"fmt"
	"lld/elevator/src/entities"
)

type CostOptimizerScheduler struct {
	name string
}

// Schedule contains the logic via which it will decide which elevator-car need to be stopped for the
// move request
func (s *CostOptimizerScheduler) Schedule(r *entities.MoveRequest) *entities.Car {
	fmt.Printf("request: %+v has been scheduled according to %v",
		*r, s.name)
	return nil
}
