package intcode

import (
	. "github.com/asuahsahua/advent2019/cmd/common"
)

// Parameter Mode Codes
const (
	PM_POINTER  = 0
	PM_IMMEDIATE = 1
	PM_RELATIVE = 2
)

type InstructionContext struct{
	// The machine processing this instruction
	Machine *IntcodeMachine
	// The ID for the function that we're executing
	FunctionID Instruction
	// CAUTION! DirectMemoryAccess to the memory in Machine!
	Parameters []*int64
	// The next instruction pointer after this instruction runs
	NextInstPtr int64
}

// -- Day 5 --
// Calculate the parameter mode of the argument in the specific position.
// With opcode 1002:
//                      ABCDE
//                       1002
//
// DE - two-digit opcode,      02 == opcode 2
//  C - mode of 1st parameter,  0 == position mode
//  B - mode of 2nd parameter,  1 == immediate mode
//  A - mode of 3rd parameter,  0 == position mode, omitted due to being a leading zero
func (m *IntcodeMachine) BuildInstructionContext(memptr int64) *InstructionContext {
	// Find the opcode, split it up
	opcode := m.Memory[memptr]
	opcodeExpanded := FillZeroes(DecimalDigitsReverse(int(opcode)), 2 + MAX_PARAMETER_COUNT)
	instruction := Instruction(10 * opcodeExpanded[1] + opcodeExpanded[0])
	paramCount := instruction.ParameterCount()
	paramModes := opcodeExpanded[2:]

	// Build parameters as pointers TODO: Abstract to another function
	params := []*int64{}
	for i := int64(0); i < paramCount; i++ {
		pvalue := &m.Memory[memptr + 1 + i]
		switch paramModes[i] {
		case PM_IMMEDIATE:
			params = append(params, pvalue)
		case PM_POINTER:
			params = append(params, &m.Memory[*pvalue])

		// Day 9: The address a relative mode parameter refers to is itself plus
		// the current relative base. When the relative base is 0, relative mode
		// parameters and position mode parameters with the same value refer to
		// the same address.
		case PM_RELATIVE:
			params = append(params, &m.Memory[m.RelativeBase + *pvalue])
		default:
			Panic("Unknown parameter mode %d", opcodeExpanded[i])
		}
	}

	return &InstructionContext{
		Machine: m,
		FunctionID: instruction,
		Parameters: params,
		NextInstPtr: memptr + 1 + paramCount,
	}
}