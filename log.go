package tools

import "log"

// LogError - Log if error is not nil
func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}
