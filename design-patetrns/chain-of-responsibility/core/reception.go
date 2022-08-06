package core

import "fmt"

type reception struct {
	next IDepartment
}

func NewReception() IDepartment {
	return &reception{next: nil}
}

func (r *reception) Execute(p *Patient) {
	if p.receptionDone {
		fmt.Println("checkup already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("Patient is in reception")
	if r.next != nil {
		r.next.Execute(p)
	}
}

func (r *reception) SetNext(department IDepartment) {
	r.next = department
}
