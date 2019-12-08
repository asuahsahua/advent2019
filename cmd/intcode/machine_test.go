package intcode

import (
	"testing"
	. "github.com/stretchr/testify/assert"
)

func TestIntcodeMachine(t *testing.T) {
	// Given example
	executeMachineTest(t, `1,9,10,3,2,3,11,0,99,30,40,50`, `3500,9,10,70,2,3,11,0,99,30,40,50`)
	// 1,0,0,0,99 becomes 2,0,0,0,99 (1 + 1 = 2).
	executeMachineTest(t, `1,0,0,0,99`, `2,0,0,0,99`)
	// 2,3,0,3,99 becomes 2,3,0,6,99 (3 * 2 = 6).
	executeMachineTest(t, `2,3,0,3,99`, `2,3,0,6,99`)
	// 2,4,4,5,99,0 becomes 2,4,4,5,99,9801 (99 * 99 = 9801).
	executeMachineTest(t, `2,4,4,5,99,0`, `2,4,4,5,99,9801`)
	// 1,1,1,4,99,5,6,0,99 becomes 30,1,1,4,2,5,6,0,99.
	executeMachineTest(t, `1,1,1,4,99,5,6,0,99`, `30,1,1,4,2,5,6,0,99`)
}

func executeMachineTest(t *testing.T, input string, expected string) {
	machine := NewIntcodeMachineStr(input)
	machine.Run()
	// Laziness: Just compare the memory of a machine loaded with our expected memory state
	expectedMachine := NewIntcodeMachineStr(expected)
	Equal(t, expectedMachine.Memory, machine.Memory)
}