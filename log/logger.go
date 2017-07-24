package log

import (
	"log"
	"os"
)

var flag = log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile;

// starts with uppercase
// this logger can be accessed from out packages
var XLog = log.New(os.Stdout, "", flag)

