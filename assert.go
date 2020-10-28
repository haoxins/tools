package tools

import (
	"log"

	"github.com/pkg/errors"
)

// AssertError - Deprecated
func AssertError(err error) {
	log.Println("Deprecated AssertError, prefer PanicError")
	if err != nil {
		log.Println(err)
		panic(errors.Wrap(err, "[Assert error failed]"))
	}
}

// PanicError panic if error not nil
func PanicError(err error) {
	if err != nil {
		log.Println(err)
		panic(errors.Wrap(err, "[Panic error]"))
	}
}

// FatalError fatal if error not nil
func FatalError(err error) {
	if err != nil {
		log.Println(err)
		log.Fatalln(errors.Wrap(err, "[Fatal error]"))
	}
}
