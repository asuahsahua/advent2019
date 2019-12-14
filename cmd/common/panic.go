package common

import (
	"fmt"
)

func Panic(format string, a ...interface{}) {
	panic(fmt.Sprintf(format, a...))
}

func PanicIf(v bool, format string, a ...interface{}) {
	if v == true {
		Panic(format, a...)
	}
}
func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}