package core

// normalBuilder builds the normal houses
type normalBuilder struct {
	house *House
}

func newNormalBuilder() IBuilder {
	return &normalBuilder{house: &House{}}
}

func (b *normalBuilder) SetHouseType() {
	b.house.HouseType = "normal"
}

func (b *normalBuilder) SetNumberOfFloors() {
	b.house.Floor = 2
}

func (b *normalBuilder) SetDoorType() {
	b.house.DoorType = "wooden door"
}

func (b *normalBuilder) SetWindowType() {
	b.house.WindowType = "wooden window"
}

func (b *normalBuilder) GetHouse() *House {
	return b.house
}
