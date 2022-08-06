package core

type IDepartment interface {
	Execute(p *Patient)
	SetNext(department IDepartment)
}
