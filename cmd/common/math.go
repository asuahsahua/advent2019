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