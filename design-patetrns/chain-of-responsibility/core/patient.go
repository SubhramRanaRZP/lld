package core

type Patient struct {
	name          string
	receptionDone bool
	checkupDone   bool
	paymentDone   bool
}

func NewPatient(name string) *Patient {
	return &Patient{
		name:          name,
		receptionDone: false,
		checkupDone:   false,
		paymentDone:   false,
	}
}
