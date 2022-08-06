package products

import (
	"fmt"
	"lld/design-patetrns/observer/core"
)

type macBookPro struct {
	name        string
	itemCount   int
	subscribers []core.ISubscriber
}

func NewMacBookPro(name string) core.IProduct {
	return &macBookPro{
		name:        name,
		itemCount:   0,
		subscribers: nil,
	}
}

func (p *macBookPro) Subscribe(subscriber core.ISubscriber) {
	p.subscribers = append(p.subscribers, subscriber)
	fmt.Printf("%v, is added to the subscriber list\n", subscriber.GetName())
}

func (p *macBookPro) UnSubscribe(subscriber core.ISubscriber) {
	for i, sub := range p.subscribers {
		if sub.GetName() == subscriber.GetName() {
			p.subscribers = append(p.subscribers[:i], p.subscribers[i+1:]...)
			break
		}
	}
	fmt.Printf("%v, is removed from the subscriber list\n", subscriber.GetName())
}

func (p *macBookPro) AddProductItem(extraProductCount int) {
	if p.itemCount == 0 {
		p.NotifySubscribers()
	}
	p.itemCount += extraProductCount
}

func (p *macBookPro) NotifySubscribers() {
	for _, sub := range p.subscribers { // notify all subscribers for this product
		sub.Notify(fmt.Sprintf("Hey %v. %v is availabel now !!!\n", sub.GetName(), p.name))
	}
}

func (p *macBookPro) MakeInventoryEmpty() {
	p.itemCount = 0
}
