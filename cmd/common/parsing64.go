package common

import (
	"strings"
	"strconv"
)

func ReadInt64s(str string) (integers []int64) {
	return SplitInt64s(str, "\n")
}

func SplitInt64s(str string, sep string) []int64 {
	integers := make([]int64, 0)
	split := strings.Split(str, sep)
	for _, v := range(split) {
		val, err := strconv.ParseInt(v, 10, 64)
		PanicIfErr(err)

		integers = append(integers, val)
	}

	return integers
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