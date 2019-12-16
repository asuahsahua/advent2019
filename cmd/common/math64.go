package common

func MinI64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func MaxI64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func AbsI64(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Returns a slice of digits representing the different decimal positions from
// most significant to least significant digit
// DecimalDigits(123456) == [1, 2, 3, 4, 5, 6]
func DecimalDigits64(value int64) []int64 {
	// Turns out doing it in reverse first is easier
	digits := DecimalDigitsReverse64(value)

	// Need to reverse the digits now
	reversed := make([]int64, len(digits))
	for i, digit := range(digits) {
		reversed[len(digits) - i - 1] = digit
	}

	return reversed
}

// Returns a slice of digits representing the different decimal positions from
// least significant to most significant digit
// DecimalDigitsReverse(123456) == [6, 5, 4, 3, 2, 1]
func DecimalDigitsReverse64(value int64) []int64 {
	digits := make([]int64, 0)
	for value > 0 {
		digits = append(digits, int64(value) % 10)
		value = value / 10
	}
	return digits
}

// Converts a string of digits to an int slice of digits
func DecimalDigitsStr64(str string) []int64 {
	digits := make([]int64, 0)
	for _, char := range str {
		PanicIf(char < '0' || char > '9', "Character string should only contain integers")
		digits = append(digits, int64(char - '0'))
	}
	return digits
}