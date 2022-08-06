package main

import (
	"lld/design-patetrns/abstract-factory/core"
	"lld/design-patetrns/abstract-factory/core/enums"
)

func main() {
	var macGUI, winGUI core.GUI

	macGUI = core.NewGUI(enums.Mac)
	winGUI = core.NewGUI(enums.Window)

	macGUI.CreateButton("red", enums.Blink)
	macGUI.ClickButton()

	winGUI.CreateButton("blue", enums.PopOut)
	winGUI.ClickButton()
}
