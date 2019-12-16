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
	// Turns out doing it in reverse first is easier
	digits := DecimalDigitsReverse(value)

	// Need to reverse the digits now
	reversed := make([]int, len(digits))
	for i, digit := range(digits) {
		reversed[len(digits) - i - 1] = digit
	}

	return reversed
}

// Returns a slice of digits representing the different decimal positions from
// least significant to most significant digit
// DecimalDigitsReverse(123456) == [6, 5, 4, 3, 2, 1]
func DecimalDigitsReverse(value int) []int {
	digits := make([]int, 0)
	for value > 0 {
		digits = append(digits, int(value) % 10)
		value = value / 10
	}
	return digits
}

// Converts a string of digits to an int slice of digits
func DecimalDigitsStr(str string) []int {
	digits := make([]int, 0)
	for _, char := range str {
		PanicIf(char < '0' || char > '9', "Character string should only contain integers")
		digits = append(digits, int(char - '0'))
	}
	return digits
}