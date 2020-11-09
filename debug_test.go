package tools

import (
	"testing"
)

func Test_Debug(t *testing.T) {
	Debug("Debug Debug Debug %d %d", 12306, 666)
}
