package main

import (
	"github.com/asuahsahua/advent2019/cmd/intcode"
	. "github.com/asuahsahua/advent2019/cmd/common"
)

func main() {
	// --- Day 5: Sunny with a Chance of Asteroids ---

	// You're starting to sweat as the ship makes its way toward Mercury. The
	// Elves suggest that you get the air conditioner working by upgrading your
	// ship computer to support the Thermal Environment Supervision Terminal.

	// The Thermal Environment Supervision Terminal (TEST) starts by running a
	// diagnostic program (your puzzle input).
	Part1("%d", RunDiagnostic())

	// --- Part Two ---

	// The air conditioner comes online! Its cold air feels good for a while,
	// but then the TEST alarms start to go off. Since the air conditioner can't
	// vent its heat anywhere but back into the spacecraft, it's actually making
	// the air inside the ship warmer.

	// Instead, you'll need to use the TEST to extend the thermal radiators.
	// This time, when the TEST diagnostic program runs its input instruction to
	// get the ID of the system to test, provide it 5, the ID for the ship's
	// thermal radiator controller. This diagnostic test suite only outputs one
	// number, the diagnostic code.

	// What is the diagnostic code for system ID 5?
	Part2("%d", intcode.RunProgram(Input, 5))
}

var Input = `3,225,1,225,6,6,1100,1,238,225,104,0,1102,46,47,225,2,122,130,224,101,-1998,224,224,4,224,1002,223,8,223,1001,224,6,224,1,224,223,223,1102,61,51,225,102,32,92,224,101,-800,224,224,4,224,1002,223,8,223,1001,224,1,224,1,223,224,223,1101,61,64,225,1001,118,25,224,101,-106,224,224,4,224,1002,223,8,223,101,1,224,224,1,224,223,223,1102,33,25,225,1102,73,67,224,101,-4891,224,224,4,224,1002,223,8,223,1001,224,4,224,1,224,223,223,1101,14,81,225,1102,17,74,225,1102,52,67,225,1101,94,27,225,101,71,39,224,101,-132,224,224,4,224,1002,223,8,223,101,5,224,224,1,224,223,223,1002,14,38,224,101,-1786,224,224,4,224,102,8,223,223,1001,224,2,224,1,223,224,223,1,65,126,224,1001,224,-128,224,4,224,1002,223,8,223,101,6,224,224,1,224,223,223,1101,81,40,224,1001,224,-121,224,4,224,102,8,223,223,101,4,224,224,1,223,224,223,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,1008,677,226,224,1002,223,2,223,1005,224,329,1001,223,1,223,107,677,677,224,102,2,223,223,1005,224,344,101,1,223,223,1107,677,677,224,102,2,223,223,1005,224,359,1001,223,1,223,1108,226,226,224,1002,223,2,223,1006,224,374,101,1,223,223,107,226,226,224,1002,223,2,223,1005,224,389,1001,223,1,223,108,226,226,224,1002,223,2,223,1005,224,404,1001,223,1,223,1008,677,677,224,1002,223,2,223,1006,224,419,1001,223,1,223,1107,677,226,224,102,2,223,223,1005,224,434,1001,223,1,223,108,226,677,224,102,2,223,223,1006,224,449,1001,223,1,223,8,677,226,224,102,2,223,223,1006,224,464,1001,223,1,223,1007,677,226,224,1002,223,2,223,1006,224,479,1001,223,1,223,1007,677,677,224,1002,223,2,223,1005,224,494,1001,223,1,223,1107,226,677,224,1002,223,2,223,1006,224,509,101,1,223,223,1108,226,677,224,102,2,223,223,1005,224,524,1001,223,1,223,7,226,226,224,102,2,223,223,1005,224,539,1001,223,1,223,8,677,677,224,1002,223,2,223,1005,224,554,101,1,223,223,107,677,226,224,102,2,223,223,1006,224,569,1001,223,1,223,7,226,677,224,1002,223,2,223,1005,224,584,1001,223,1,223,1008,226,226,224,1002,223,2,223,1006,224,599,101,1,223,223,1108,677,226,224,102,2,223,223,1006,224,614,101,1,223,223,7,677,226,224,102,2,223,223,1005,224,629,1001,223,1,223,8,226,677,224,1002,223,2,223,1006,224,644,101,1,223,223,1007,226,226,224,102,2,223,223,1005,224,659,101,1,223,223,108,677,677,224,1002,223,2,223,1006,224,674,1001,223,1,223,4,223,99,226`

func RunDiagnostic() int64 {
	machine := intcode.NewIntcodeMachineStr(Input)

	// The TEST diagnostic program will start by requesting from the user the ID
	// of the system to test by running an input instruction - provide it 1, the
	// ID for the ship's air conditioner unit.
	machine.Input <- 1
	
	// It will then perform a series of diagnostic tests confirming that various
	// parts of the Intcode computer, like parameter modes, function correctly.

	// For each test, it will run an output instruction indicating how far the
	// result of the test was from the expected value, where 0 means the test
	// was successful. Non-zero outputs mean that a function is not working
	// correctly; check the instructions that were run before the output
	// instruction to see which one failed.
	go machine.Run()
	// read until its not zero? this doesn't seem deterministic  enough, lol
	var out int64 = 0
	for out == 0  {
		out = <- machine.Output
	}
	
	// Finally, the program will output a diagnostic code and immediately halt.
	// This final output isn't an error; an output followed immediately by a
	// halt means the program finished. If all outputs were zero except the
	// diagnostic code, the diagnostic program ran successfully.

	// After providing 1 to the only input instruction 
	// and passing all the tests, what diagnostic code does the program produce?
	return out
}