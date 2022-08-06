package main

import (
	"lld/design-patetrns/chain-of-responsibility/core"
)

func main() {
	p := core.NewPatient("Shyam")

	var r, d, py core.IDepartment

	r = core.NewReception()
	d = core.NewDoctor()
	py = core.NewPayment()

	d.SetNext(py)
	r.SetNext(d)

	r.Execute(p)
}
