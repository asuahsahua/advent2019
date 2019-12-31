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

// DecimalDigits returns a slice of digits representing the different decimal
// positions from most significant to least significant digit
// DecimalDigits(123456) == [1, 2, 3, 4, 5, 6]
func DecimalDigits(value int) []int {
	// Turns out doing it in reverse first is easier
	digits := DecimalDigitsReverse(value)

	// Need to reverse the digits now
	reversed := make([]int, len(digits))
	for i, digit := range digits {
		reversed[len(digits)-i-1] = digit
	}

	return reversed
}

// DecimalDigitsReverse returns a slice of digits representing the different
// decimal positions from least significant to most significant digit
// DecimalDigitsReverse(123456) == [6, 5, 4, 3, 2, 1]
func DecimalDigitsReverse(value int) []int {
	digits := make([]int, 0)
	for value > 0 {
		digits = append(digits, int(value)%10)
		value = value / 10
	}
	return digits
}

// DecimalDigitsStr converts a string of digits to an int slice of digits
func DecimalDigitsStr(str string) []int {
	digits := make([]int, 0)
	for _, char := range str {
		PanicIf(char < '0' || char > '9', "Character string should only contain integers")
		digits = append(digits, int(char-'0'))
	}
	return digits
}

// GCD finds the greatest common divisor using the Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM finds the Least Common Multiple of many integers
func LCM(a, b int, ints ...int) int {
	// For a, b ∈ ℕ, a*b = LCM(a, b)*GCD(a, b)
	lcm := (a * b) / GCD(a, b)

	for _, c := range ints {
		lcm = LCM(lcm, c)
	}

	return lcm
}

// CeilDiv divides the numerator by the denominator and returns the ceiling
func CeilDiv(num, denom int) int {
	div := num / denom
	mod := num % denom
	if mod > 0 {
		return div + 1
	}
	return div
}

func SumInts(ints []int) int {
	sum := 0
	for _, val := range ints {
		sum += val
	}
	return sum
}