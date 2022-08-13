package entities

import (
	"fmt"
	"lld/elevator/src/enums"
)

type Door struct {
	id       string
	doorType enums.DoorType
}

func (d *Door) Open() {
	fmt.Printf("door id: %v of type %v opened", d.id, d.doorType)
}

func (d *Door) Close() {
	fmt.Printf("door id: %v of type %v closed", d.id, d.doorType)
}
