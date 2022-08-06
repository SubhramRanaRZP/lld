package core

type IBuilder interface {
	SetHouseType()
	SetNumberOfFloors()
	SetDoorType()
	SetWindowType()
	GetHouse() *House
}

type IDirector interface {
	SetBuilder(builder IBuilder)
	BuildHouse()
	ShowHouse()
}
