package core

type GUI interface {
	CreateButton(buttonType string)
	ClickButton()
}

type Button interface {
	Click()
}
