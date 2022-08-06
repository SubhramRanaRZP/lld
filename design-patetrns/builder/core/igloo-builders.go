package core

import (
	"lld/design-patetrns/builder/core/enums"
)

func NewBuilder(houseType enums.HouseType) IBuilder {
	var b IBuilder
	switch houseType {
	case enums.Igloo:
		b = newIglooBuilder()
	case enums.Normal:
		b = newNormalBuilder()
	default:
		b = nil
	}
	return b
}

// iglooBuilders builds the igloo houses
type iglooBuilders struct {
	house *House
}

func newIglooBuilder() IBuilder {
	return &iglooBuilders{house: &House{}}
}

func (b *iglooBuilders) SetNumberOfFloors() {
	b.house.Floor = 1
}

func (b *iglooBuilders) SetDoorType() {
	b.house.DoorType = "snow door"
}

func (b *iglooBuilders) SetWindowType() {
	b.house.WindowType = "snow window"
}

func (b *iglooBuilders) GetHouse() *House {
	return b.house
}
