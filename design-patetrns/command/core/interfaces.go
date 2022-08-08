package core

type IDevice interface {
	On()
}

type ICommand interface {
	Execute()
}

type IButton interface {
	Click()
}
