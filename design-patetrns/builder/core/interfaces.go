package core

type IBuilder interface {
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
