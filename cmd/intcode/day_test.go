package intcode

import (
	"testing"
	. "github.com/stretchr/testify/assert"
)

// This file holds nonspecific tests from each day's prompts

// --- Day 5 - Part 2 ---
func TestDay05Part2(t *testing.T) {
	// Using position mode, consider whether the input is equal to 8; output 1
	// (if it is) or 0 (if it is not).
	prog1 := `3,9,8,9,10,9,4,9,99,-1,8`
	Equal(t, 1, RunProgram(prog1, 8))
	Equal(t, 0, RunProgram(prog1, 7))
	Equal(t, 0, RunProgram(prog1, 1))

	// Using position mode, consider whether the input is less than 8; output 1
	// (if it is) or 0 (if it is not).
	prog2 := `3,9,7,9,10,9,4,9,99,-1,8`
	Equal(t, 1, RunProgram(prog2, 3))
	Equal(t, 1, RunProgram(prog2, 7))
	Equal(t, 0, RunProgram(prog2, 8))
	Equal(t, 0, RunProgram(prog2, 15))
	
	// Using immediate mode, consider whether the input is equal to 8; output 1
	// (if it is) or 0 (if it is not).
	prog3 := `3,3,1108,-1,8,3,4,3,99`
	Equal(t, 1, RunProgram(prog3, 8))
	Equal(t, 0, RunProgram(prog3, 7))
	Equal(t, 0, RunProgram(prog3, 15))

	// Using immediate mode, consider whether the input is less than 8; output 1
	// (if it is) or 0 (if it is not).
	prog4 := `3,3,1107,-1,8,3,4,3,99`
	Equal(t, 1, RunProgram(prog4, 3))
	Equal(t, 1, RunProgram(prog4, 7))
	Equal(t, 0, RunProgram(prog4, 8))
	Equal(t, 0, RunProgram(prog4, 15))

	// Here are some jump tests that take an input, then output 0 if the input
	// was zero or 1 if the input was non-zero:
	// using position mode
	prog5 := `3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9`
	Equal(t, 0, RunProgram(prog5, 0))
	Equal(t, 1, RunProgram(prog5, 1))
	Equal(t, 1, RunProgram(prog5, -1))
	// using immediate mode
	prog6 := `3,3,1105,-1,9,1101,0,0,12,4,12,99,1`
	Equal(t, 0, RunProgram(prog6, 0))
	Equal(t, 1, RunProgram(prog6, 1))
	Equal(t, 1, RunProgram(prog6, -1))

	// Here's a larger example:
	prog7 := `3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99`
	// The above example program uses an input instruction to ask for a single
	// number. The program will then output 999 if the input value is below 8,
	// output 1000 if the input value is equal to 8, or output 1001 if the input
	// value is greater than 8.
	Equal(t, 999, RunProgram(prog7, 7))
	Equal(t, 1000, RunProgram(prog7, 8))
	Equal(t, 1001, RunProgram(prog7, 9))
}