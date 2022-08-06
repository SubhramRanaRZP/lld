package enums

type (
	GUIType    int16
	ButtonType int16
)

const (
	Mac guiType = iota + 1
	Window
)

const (
	PopOut buttonType = iota + 1
	Blink
)
