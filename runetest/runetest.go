package main

import (
	"fmt"
)

func main() {
	str := "Have a nice day, 中国!"
	r := []rune(str)
	for idx, element := range r {
		// convert rune to string for display using string()
		fmt.Printf("element at %d is %v\n", idx, string(element))
	}
	fmt.Println([]rune("中")[0])
}
