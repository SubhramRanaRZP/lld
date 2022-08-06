package core

import "fmt"

type payment struct {
	next IDepartment
}

func NewPayment() IDepartment {
	return &payment{next: nil}
}

func (py *payment) Execute(p *Patient) {
	if p.receptionDone {
		fmt.Println("payment already done")
		py.next.Execute(p)
		return
	}

	fmt.Println("Patient is doing payment processing")

	if py.next != nil {
		py.next.Execute(p)
	}
}

func (py *payment) SetNext(department IDepartment) {
	py.next = department
}
