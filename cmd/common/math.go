package common

func MinI(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxI(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func AbsI(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Returns a slice of digits representing the different decimal positions from
// most significant to least significant digit
// DecimalDigits(123456) == [1, 2, 3, 4, 5, 6]
func DecimalDigits(value int) []int {
	digits := make([]int, 0)
	for value > 0 {
		digits = append(digits, value % 10)
		value = value / 10
	}

	// Need to reverse the digits now
	reversed := make([]int, len(digits))
	for i, digit := range(digits) {
		reversed[len(digits) - i - 1] = digit
	}

	return reversed
}