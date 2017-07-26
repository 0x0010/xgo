package main

import (
	"github.com/0x0010/xgo/xtype"
	"fmt"
)

func main() {
	android := new(xtype.Android)
	android.SetName("Helen")
	android.Talk()

	android2 := xtype.NewAndroid("Jobs", "model")
	fmt.Println(android2)
	android2.Talk()
}
