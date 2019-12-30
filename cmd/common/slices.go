package common

import (
	"strconv"
	"strings"
)

func FillZeroes(slice []int, count int) []int {
	// copy the slice to begin with so we don't mess with the underlying data structure
	out := make([]int, len(slice))
	copy(out, slice)

	for i := len(slice); i < count; i++ {
		out = append(out, 0)
	}

	return out
}

// Fills an array of count size with repeats of slice
func FillRepeat(slice []int, count int) []int {
	out := make([]int, count)
	for i := 0; i < count; i++ {
		out[i] = slice[i % len(slice)]
	}
	return out
}

func JoinI(slice []int, sep string) string {
	strs := make([]string, 0)
	for _, v := range slice {
		strs = append(strs, strconv.Itoa(v))
	}
	return strings.Join(strs, sep)
}