package main

import (
	"github.com/0x0010/xgo/xtype"
)

func main() {
	jack := new(xtype.Person)
	jack.SetName("Jack")
	jack.Talk()
}
