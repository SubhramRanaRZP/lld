package entities

import "fmt"

type Alarm struct {
	ringTune string
}

func (a *Alarm) Ring() {
	fmt.Printf("alarm ringing... %v", a.ringTune)
}
