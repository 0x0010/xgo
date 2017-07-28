package main

import (
	"fmt"
	"github.com/0x0010/xgo/utils"
)

func main() {
	f := utils.Fibonacci()
	for i := 0; i < 1000; i++ {
		fmt.Printf("F(%v):%v\n",i, f())
	}
}
