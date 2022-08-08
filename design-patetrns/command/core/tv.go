package core

import "fmt"

type tv struct{}

func NewTv() IDevice {
	return &tv{}
}

func (t *tv) On() {
	fmt.Println("TV is on")
}
