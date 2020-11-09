package tools

import (
	"fmt"
	"os"
)

var (
	enableDebug = os.Getenv("GO_DEBUG") != ""
)

// Debug Just a alias of fmt.Printf with a ENV flag
func Debug(format string, a ...interface{}) {
	if enableDebug {
		fmt.Printf("[Debug] "+format+"\n", a...)
	}
}
