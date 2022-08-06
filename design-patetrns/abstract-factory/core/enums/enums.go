package enums

type (
	GUIType    int16
	ButtonType int16
)

const (
	Mac GUIType = iota + 1
	Window
)

const (
	PopOut ButtonType = iota + 1
	Blink
)
