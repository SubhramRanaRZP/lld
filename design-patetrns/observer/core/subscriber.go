package core

import "fmt"

type Customer struct {
	name string
}

func NewCustomer(name string) *Customer {
	return &Customer{name: name}
}

func (s *Customer) GetName() string {
	return s.name
}

func (s *Customer) Notify(msg string) {
	// this can be implemented by strategy pattern.
	// this can be easily extensible for gmail, ph no, sms etc...
	// for implementation purpose, I am just printing here
	fmt.Println(msg)
}
