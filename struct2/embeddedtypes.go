package main

import (
	"github.com/0x0010/xgo/xpkg"
)

func main() {
	android := new(xpkg.Android)
	android.SetName("Helen")
	android.Talk()
}

