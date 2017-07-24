package main

import (
	"github.com/0x0010/xgo/xtype"
)

func main() {
	android := new(xtype.Android)
	android.SetName("Helen")
	android.Talk()

	android = xtype.NewAndroid("Jobs", "model")
	android.Talk()
}
