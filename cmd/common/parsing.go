package common

import (
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

func CommaSeparatedToInt64(str string) (integers []int64) {
	splitUp := strings.Split(str, ",")
	for _, v := range(splitUp) {
		val, err := strconv.ParseInt(v, 10, 64)
		PanicIfErr(err)

		integers = append(integers, val)
	}

	return
}