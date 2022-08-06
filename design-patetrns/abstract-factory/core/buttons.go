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
	fmt.Printf("%v colored button poped-out\n", b.color)
}

// blinkButton blinks when you click
type blinkButton struct {
	color string
}

func (b *blinkButton) Click() {
	fmt.Printf("%v colored button blinked\n", b.color)
}

func newButton(color string, buttonType enums.ButtonType) Button {
	var b Button
	switch buttonType {
	case enums.PopOut:
		b = &popButton{color: color}
	case enums.Blink:
		b = &blinkButton{color: color}
	default:
		b = nil
	}
	return b
}
