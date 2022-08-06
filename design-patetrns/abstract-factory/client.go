package main

import "lld/design-patetrns/abstract-factory/core"

func main() {
	var macGUI, winGUI core.GUI

	macGUI = core.NewGUI("mac")
	winGUI = core.NewGUI("win")

	macGUI.ClickButton()
	winGUI.ClickButton()
}
