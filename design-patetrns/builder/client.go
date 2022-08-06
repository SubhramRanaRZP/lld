package main

import (
	"lld/design-patetrns/builder/core"
	"lld/design-patetrns/builder/core/enums"
)

func main() {
	iglooBuilder := core.NewBuilder(enums.Igloo)

	director := core.NewDirector(iglooBuilder)
	director.BuildHouse()
	director.ShowHouse()

	// change the builder type
	normalBuilder := core.NewBuilder(enums.Normal)
	director.SetBuilder(normalBuilder)
	director.BuildHouse()
	director.ShowHouse()
}
