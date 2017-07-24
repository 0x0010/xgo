package main

import (
	"github.com/0x0010/xgo/xtype"
)

func main() {
	jack := xtype.NewPerson("Jack[Set by constructor]")
	jack.Talk()
	jack.SetName("Jack[Set by SetName method]")
	jack.Talk()
}
