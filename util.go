package tools

import (
	"fmt"
	"log"
)

// RightPad right pad
func RightPad(content string, length string) string {
	log.Println("Deprecated RightPad, prefer github.com/willf/pad")
	return fmt.Sprintf("%-"+length+"v", content)
}
