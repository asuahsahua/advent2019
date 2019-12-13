package common

import (
	"fmt"
)

func Part1(format string, a ...interface{}) {
	fmt.Printf(
		"Part 1: %s\n",
		fmt.Sprintf(format, a...),
	)
}

func Part2(format string, a ...interface{}) {
	fmt.Printf(
		"Part 2: %s\n",
		fmt.Sprintf(format, a...),
	)
}

func Debug(format string, a ...interface{}) {
	fmt.Printf(
		"DEBUG: %s\n",
		fmt.Sprintf(format, a...),
	)
}