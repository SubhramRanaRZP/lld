package core

type IProduct interface {
	Subscribe(subscriber ISubscriber)
	UnSubscribe(subscriber ISubscriber)
	AddProductItem(extraProductCount int)
	NotifySubscribers()
	MakeInventoryEmpty()
}

type ISubscriber interface {
	GetName() string
	Notify(msg string) // we can in turn implement this by strategy pattern
}
