package intcode

import (
	. "github.com/asuahsahua/advent2019/cmd/common"
)

const MAX_PARAMETER_COUNT = 3

func (machine *IntcodeMachine) ParameterCount(functionID int) int {
	switch functionID {
	case INST_ADD:
		return 3
	case INST_MULT:
		return 3
	case INST_INPUT:
		return 1
	case INST_OUTPUT:
		return 1
	case INST_HCF:
		return 0
	default:
		Panic("could not determine parameter count for function %d", functionID)
	}

	return -1
}

func (ctx InstructionContext) Execute() {
	switch ctx.FunctionID {
	case INST_ADD:
		ctx.I01_Add()
	case INST_MULT:
		ctx.I02_Mult()
	case INST_INPUT:
		ctx.I03_Input()
	case INST_OUTPUT:
		ctx.I04_Output()
	case INST_HCF:
		ctx.I99_HCF()
	}
}