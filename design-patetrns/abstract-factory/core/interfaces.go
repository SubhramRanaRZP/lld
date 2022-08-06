package core

import "lld/design-patetrns/abstract-factory/core/enums"

type GUI interface {
	CreateButton(color string , buttonType enums.ButtonType)
	ClickButton()
}

type Button interface {
	Click()
}
