package main

import (
	. "github.com/asuahsahua/advent2019/cmd/intcode"
	. "github.com/asuahsahua/advent2019/cmd/common"
)

func main() {
	// --- Day 2: 1202 Program Alarm ---

	// On the way to your gravity assist around the Moon, your ship computer
	// beeps angrily about a "1202 program alarm". On the radio, an Elf is
	// already explaining how to handle the situation: "Don't worry, that's
	// perfectly norma--" The ship computer bursts into flames.
	
	// You notify the Elves that the computer's magic smoke seems to have
	// escaped. "That computer ran Intcode programs like the gravity assist
	// program it was working on; surely there are enough spare parts up there
	// to build a new Intcode computer!"
	
	// Once you have a working computer, the first step is to restore the
	// gravity assist program (your puzzle input) to the "1202 program alarm"
	// state it had just before the last computer caught fire. To do this,
	// before running the program, replace position 1 with the value 12 and
	// replace position 2 with the value 2. 

	// -- Refactored a bit for Part 2

	// >>> What value is left at position 0 after the program halts?
	Part1("%d", runProgram(Input1, 12, 2))

	// 	--- Part Two ---

	// "Good, the new computer seems to be working correctly! Keep it nearby
	// during this mission - you'll probably use it again. Real Intcode
	// computers support many more features than your new one, but we'll let you
	// know what they are as you need them."
	
	// "However, your current priority should be to complete your gravity assist
	// around the Moon. For this mission to succeed, we should settle on some
	// terminology for the parts you've already built."
	
	// "With terminology out of the way, we're ready to proceed. To complete the
	// gravity assist, you need to determine what pair of inputs produces the
	// output 19690720."
	requiredOutput := 19690720
	
	// Find the input noun and verb that cause the program to produce the output
	// 19690720. What is 100 * noun + verb? (For example, if noun=12 and verb=2,
	// the answer would be 1202.)
	noun, verb := findRequiredOutput(Input1, requiredOutput)
	Part2("%d", 100 * noun + verb)
}

func runProgram(program string, noun int, verb int) int {
	machine := NewIntcodeMachineStr(program)

	// The inputs should still be provided to the program by replacing the
	// values at addresses 1 and 2, just like before. In this program, the value
	// placed in address 1 is called the noun, and the value placed in address 2
	// is called the verb. Each of the two input values will be between 0 and
	// 99, inclusive.
	// TODO: Validate between 0 and 99
	machine.Memory[1] = int64(noun)
	machine.Memory[2] = int64(verb)

	//What value is left at position 0 after the program halts?
	machine.Run()

	return int(machine.Memory[0])
}

func findRequiredOutput(program string, requiredOutput int) (noun int, verb int) {
	for noun = 0; noun <= 100; noun++ {
		for verb = 0; verb <= 100; verb++ {
			if requiredOutput == runProgram(Input1, noun, verb) {
				return
			}
		}
	}
	panic("Could not find the required output value :(")
}

var Input1 string = `1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,9,1,19,1,9,19,23,1,23,5,27,2,27,10,31,1,6,31,35,1,6,35,39,2,9,39,43,1,6,43,47,1,47,5,51,1,51,13,55,1,55,13,59,1,59,5,63,2,63,6,67,1,5,67,71,1,71,13,75,1,10,75,79,2,79,6,83,2,9,83,87,1,5,87,91,1,91,5,95,2,9,95,99,1,6,99,103,1,9,103,107,2,9,107,111,1,111,6,115,2,9,115,119,1,119,6,123,1,123,9,127,2,127,13,131,1,131,9,135,1,10,135,139,2,139,10,143,1,143,5,147,2,147,6,151,1,151,5,155,1,2,155,159,1,6,159,0,99,2,0,14,0`