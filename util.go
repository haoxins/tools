package tools

import (
	"fmt"
)

// RightPad right pad
// Deprecated - github.com/willf/pad
func RightPad(content string, length string) string {
	return fmt.Sprintf("%-"+length+"v", content)
}
