package intcode

import (
	"testing"
	. "github.com/stretchr/testify/assert"
)

func TestInstructionContextOpcode(t *testing.T) {
	machine := NewIntcodeMachineStr(`1002,4,3,4,33`)

	// Build a context around the 0th command
	ctx := machine.BuildInstructionContext(0)

	Equal(t, []int{1002, 4, 3, 4, 33}, machine.Memory)

	// Do some asserts
	Equal(t, 2, ctx.FunctionID)
	Equal(t, 33, *ctx.Parameters[0])
	Equal(t, 3, *ctx.Parameters[1])
	Equal(t, 33, *ctx.Parameters[2])

	ctx.Execute()

	Equal(t, 99, machine.Memory[4])
}