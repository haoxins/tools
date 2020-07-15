package tools

import (
	"log"

	"github.com/pkg/errors"
)

// AssertError - Panic if error is not nil
func AssertError(err error) {
	if err != nil {
		log.Fatalln(err)
		panic(errors.Wrap(err, "[Assert error failed]"))
	}
}
