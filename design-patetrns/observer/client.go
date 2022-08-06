package main

import (
	"lld/design-patetrns/observer/core"
	"lld/design-patetrns/observer/core/products"
)

func main() {
	ram := core.NewCustomer("Ram")
	shyam := core.NewCustomer("Shyam")

	macBookPro := products.NewMacBookPro("macBookPro - 16'")

	macBookPro.Subscribe(ram)

	macBookPro.AddProductItem(3) // only Ram will get notified
	macBookPro.AddProductItem(1) // no one will get notified

	macBookPro.MakeInventoryEmpty()

	macBookPro.Subscribe(shyam)

	macBookPro.AddProductItem(3) // both Ram and Shyam will get notified

	macBookPro.MakeInventoryEmpty()

	macBookPro.UnSubscribe(ram)

	macBookPro.AddProductItem(1) // only Shyam get notified
}
