package intcode

import (
	"testing"
	. "github.com/stretchr/testify/assert"
)

// This file holds nonspecific tests from each day's prompts

// Day 5 - Part 2
func TestDay05Part2(t *testing.T) {
	// Using position mode, consider whether the input is equal to 8; output 1
	// (if it is) or 0 (if it is not).
	prog1 := `3,9,8,9,10,9,4,9,99,-1,8`
	Equal(t, int64(1), RunProgram(prog1, 8))
	Equal(t, int64(0), RunProgram(prog1, 7))
	Equal(t, int64(0), RunProgram(prog1, 1))

	// Using position mode, consider whether the input is less than 8; output 1
	// (if it is) or 0 (if it is not).
	prog2 := `3,9,7,9,10,9,4,9,99,-1,8`
	Equal(t, int64(1), RunProgram(prog2, 3))
	Equal(t, int64(1), RunProgram(prog2, 7))
	Equal(t, int64(0), RunProgram(prog2, 8))
	Equal(t, int64(0), RunProgram(prog2, 15))
	
	// Using immediate mode, consider whether the input is equal to 8; output 1
	// (if it is) or 0 (if it is not).
	prog3 := `3,3,1108,-1,8,3,4,3,99`
	Equal(t, int64(1), RunProgram(prog3, 8))
	Equal(t, int64(0), RunProgram(prog3, 7))
	Equal(t, int64(0), RunProgram(prog3, 15))

	// Using immediate mode, consider whether the input is less than 8; output 1
	// (if it is) or 0 (if it is not).
	prog4 := `3,3,1107,-1,8,3,4,3,99`
	Equal(t, int64(1), RunProgram(prog4, 3))
	Equal(t, int64(1), RunProgram(prog4, 7))
	Equal(t, int64(0), RunProgram(prog4, 8))
	Equal(t, int64(0), RunProgram(prog4, 15))

	// Here are some jump tests that take an input, then output 0 if the input
	// was zero or 1 if the input was non-zero:
	// using position mode
	prog5 := `3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9`
	Equal(t, int64(0), RunProgram(prog5, 0))
	Equal(t, int64(1), RunProgram(prog5, 1))
	Equal(t, int64(1), RunProgram(prog5, -1))
	// using immediate mode
	prog6 := `3,3,1105,-1,9,1101,0,0,12,4,12,99,1`
	Equal(t, int64(0), RunProgram(prog6, 0))
	Equal(t, int64(1), RunProgram(prog6, 1))
	Equal(t, int64(1), RunProgram(prog6, -1))

	// Here's a larger example:
	prog7 := `3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99`
	// The above example program uses an input instruction to ask for a single
	// number. The program will then output 999 if the input value is below 8,
	// output 1000 if the input value is equal to 8, or output 1001 if the input
	// value is greater than 8.
	Equal(t, int64(999), RunProgram(prog7, 7))
	Equal(t, int64(1000), RunProgram(prog7, 8))
	Equal(t, int64(1001), RunProgram(prog7, 9))
}

// Day 09 - Part 1
func TestDay09Part1(t *testing.T) {
	// Here are some example programs that use these features:
	// takes no input and produces a copy of itself as output.
	// prog1 := `109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99`

    // 1102,34915192,34915192,7,4,7,99,0 should output a 16-digit number.
    // 104,1125899906842624,99 should output the large number in the middle.
}