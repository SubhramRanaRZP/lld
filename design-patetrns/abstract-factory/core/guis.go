package core

import "lld/design-patetrns/abstract-factory/core/enums"

// macGUI is a GUI that follows MAC like button and behaviours
type macGUI struct {
	b Button
}

func (gui *macGUI) CreateButton(color string , buttonType enums.ButtonType) {
	gui.b = newButton(color, buttonType)
}

func (gui *macGUI) ClickButton() {
	gui.b.Click()
}

// winGUI is a GUI that follows MAC like button and behaviours
type winGUI struct {
	b Button
}

func (gui *winGUI) CreateButton(color string, buttonType enums.ButtonType) {
	gui.b = newButton(color, buttonType)
}

func (gui *winGUI) ClickButton() {
	gui.b.Click()
}

func NewGUI(GUIType enums.GUIType) GUI {
	var gui GUI
	switch GUIType {
	case enums.Mac:
		gui = &macGUI{}
	case enums.Window:
		gui = &winGUI{}
	}
	return gui
}
