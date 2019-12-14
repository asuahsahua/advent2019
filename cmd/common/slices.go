package common

func FillZeroes(slice []int, count int) []int {
	// copy the slice to begin with so we don't mess with the underlying data structure
	out := make([]int, len(slice))
	copy(out, slice)

	for i := len(slice); i < count; i++ {
		out = append(out, 0)
	}

	return out
}