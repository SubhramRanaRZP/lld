package core

import "fmt"

type doctor struct {
	next IDepartment
}

func NewDoctor() IDepartment {
	return &doctor{next: nil}
}

func (d *doctor) Execute(p *Patient) {
	if p.receptionDone {
		fmt.Println("checkup already done")
		d.next.Execute(p)
		return
	}
	fmt.Println("doctor check up is going on")
	if d.next != nil {
		d.next.Execute(p)
	}
}

func (d *doctor) SetNext(department IDepartment) {
	d.next = department
}
