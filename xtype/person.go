package xtype

import "fmt"

type Person struct {
	name string
}

// set person name
func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.name)
}

// person constructor
func NewPerson(name string) *Person {
	p := new(Person)
	p.name = name
	return p
}
