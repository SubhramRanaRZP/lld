package core

import (
	"encoding/json"
	"fmt"
	"log"
)

// director do the sequencing of all the required construction tasks
// it also facilitate client to select the type of builder at runtime
type director struct {
	builder IBuilder
}

func NewDirector(builder IBuilder) IDirector {
	d := director{builder: builder}
	return &d
}

func (d *director) SetBuilder(builder IBuilder) {
	d.builder = builder
}

func (d *director) BuildHouse() {
	d.builder.SetHouseType()
	d.builder.SetNumberOfFloors()
	d.builder.SetDoorType()
	d.builder.SetWindowType()
}

func (d *director) ShowHouse() {
	h := d.builder.GetHouse()
	empJSON, err := json.MarshalIndent(h, "", "   ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("House details:\n %s\n", string(empJSON))
}
