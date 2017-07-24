package xpkg

import "fmt"

type Person struct {
	name string
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.name)
}