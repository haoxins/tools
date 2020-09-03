package tools

import (
	"fmt"
)

// RightPad right pad
func RightPad(content string, length string) string {
	return fmt.Sprintf("%-"+length+"v", content)
}
