package main

import (
	"fmt"

	"github.com/0x0010/xgo/log"
	"github.com/0x0010/xgo/stringutil"
	"encoding/hex"
)

func main() {
	fmt.Printf("Hello, world.\n")
	fmt.Println(stringutil.Reverse("!oG olleH :gnirts desreveR"))
	fmt.Println(Reverse("!oG olleH :litu egakcap evitan gnisu gnirts desreveR"))
	fmt.Println(stringutil.Split("Split string from split function"))
	log.XLog.Println("Hello from xlog")

	log.XLog.Println(byte('a'))
	log.XLog.Println(hex.EncodeToString([]byte("a")))
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
