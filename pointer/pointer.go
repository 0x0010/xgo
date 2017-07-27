package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	s := "Hello World!"
	s2 := "Hello world!"
	([]rune(s2))[0] = rune('a')
	fmt.Println(strings.EqualFold(s, s2))
	fmt.Printf("Pointer type of string is %T\n", &s)
	i := 43
	fmt.Printf("Pointer type of string is %T\n", &i)
	fmt.Printf("Value of i is %v\n", *&i)

	v := Vertex{1, 2}
	v.X = 100
	vp := &v
	fmt.Println(vp.X)
	fmt.Println(v)
}
