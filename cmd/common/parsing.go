package common

import (
	"regexp"
	"strings"
	"strconv"
)

func ReadInts(str string) (integers []int) {
	return SplitInts(str, "\n")
}

func SplitLines(str string) []string {
	return strings.Split(str, "\n")
}

func SplitInts(str string, sep string) []int {
	integers := make([]int, 0)
	split := strings.Split(str, sep)
	for _, v := range(split) {
		val, err := strconv.Atoi(v)
		PanicIfErr(err)

		integers = append(integers, val)
	}

	return integers
}

func CommaSeparatedToInt(str string) (integers []int) {
	return SplitInts(str, ",")
}

// Lazy converter to grab the first int looking thing in the string
// No guarantees! Panics a lot!
var lazyIntRex = regexp.MustCompile(`(-?\d+)`)
func LazyInt(str string) int {
	val, err := strconv.Atoi(lazyIntRex.FindString(str))
	PanicIfErr(err)
	return val
}