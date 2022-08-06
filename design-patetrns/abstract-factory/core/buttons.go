package core

import (
	"fmt"
	"lld/design-patetrns/abstract-factory/core/enums"
)

// popButton pop out when you click
type popButton struct {
	color string
}

func (b *popButton) Click() {
	fmt.Printf("%v colored button poped-out", b.color)
}

// blinkButton blinks when you click
type blinkButton struct {
	color string
}

func (b *blinkButton) Click() {
	fmt.Printf("%v colored button blinked", b.color)
}

func newButton(buttonType enums.ButtonType) Button {
	var b Button
	switch buttonType {
	case enums.PopOut:
		b = &popButton{}
	case enums.Blink:
		b = &blinkButton{}
	default:
		b = nil
	}
	return b
}
