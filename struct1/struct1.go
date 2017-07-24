package main

import "fmt"

func main() {
	jack := Person{Name: "Jack"}
	jack.Talk()
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}
