package tools

import "fmt"

// Recover recover with print error
func Recover() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
