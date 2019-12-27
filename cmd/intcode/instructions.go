package intcode

import (
	"github.com/asuahsahua/advent2019/cmd/common"
)

const MAX_PARAMETER_COUNT = 3

type Instruction int64
const (
	Add Instruction = 1
	Multiply Instruction = 2
	Input Instruction = 3
	Output Instruction = 4
	JumpIfTrue Instruction = 5
	JumpIfFalse Instruction = 6
	LessThan Instruction = 7
	Equals Instruction = 8
	AdjustRelativeBase Instruction = 9
	HaltCatchFire Instruction = 99 
)

func (ic Instruction) ParameterCount() int64 {
	switch ic {
	case Add:
		return 3
	case Multiply:
		return 3
	case Input:
		return 1
	case Output:
		return 1
	case JumpIfTrue:
		return 2
	case JumpIfFalse:
		return 2
	case LessThan:
		return 3
	case Equals:
		return 3
	case AdjustRelativeBase:
		return 1
	case HaltCatchFire:
		return 0
	default:
		common.Panic("Unexpected instruction code %d", ic)
		return -1
	}
}

func (ic Instruction) Execute(ctx *InstructionContext) {
	switch ic {
	case Add:
		I01_Add(ctx)
	case Multiply:
		I02_Mult(ctx)
	case Input:
		I03_Input(ctx)
	case Output:
		I04_Output(ctx)
	case JumpIfTrue:
		I05_JumpIfTrue(ctx)
	case JumpIfFalse:
		I06_JumpIfFalse(ctx)
	case LessThan:
		I07_LessThan(ctx)
	case Equals:
		I08_Equals(ctx)
	case AdjustRelativeBase:
		I09_AdjustRelativeBase(ctx)
	case HaltCatchFire:
		I99_HCF(ctx)
	default:
		common.Panic("Unexpected instruction code %d", ic)
	}
}

func (ctx *InstructionContext) Execute() {
	ctx.FunctionID.Execute(ctx)
}

// --- Day 2 ---
// Opcode 1 adds together numbers read from two positions and stores the result
// in a third position. The three integers immediately after the opcode tell you
// these three positions - the first two indicate the positions from which you
// should read the input values, and the third indicates the position at which
// the output should be stored.
func I01_Add(ctx *InstructionContext) {
	p := ctx.Parameters
	*p[2] = *p[0] + *p[1]
}

// --- Day 2 ---
// Opcode 2 works exactly like opcode 1, except it multiplies the two inputs
// instead of adding them. Again, the three integers after the opcode indicate
// where the inputs and outputs are, not their values.
func I02_Mult(ctx *InstructionContext) {
	p := ctx.Parameters
	*p[2] = *p[0] * *p[1]
}

// --- Day 5 ---
// Opcode 3 takes a single integer as input and saves it to the position
// given by its only parameter. For example, the instruction 3,50 would take
// an input value and store it at address 50.
func I03_Input(ctx *InstructionContext)  {
	for {
		select {
		case input := <- ctx.Machine.Input:
			*ctx.Parameters[0] = input
			return
		case interrupt := <- ctx.Machine.Interrupt:
			ctx.Machine.handleInterrupt(interrupt)
		}
	}
}

// --- Day 5 ---
// Opcode 4 outputs the value of its only parameter. For example, the
// instruction 4,50 would output the value at address 50.
func I04_Output(ctx *InstructionContext) {
	for {
		select {
		case ctx.Machine.Output <- *ctx.Parameters[0]:
			return
		case interrupt := <- ctx.Machine.Interrupt:
			ctx.Machine.handleInterrupt(interrupt)
		}
	}
}

// --- Day 5 Part 2 ---
// Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the
// instruction pointer to the value from the second parameter. Otherwise, it
// does nothing.
func I05_JumpIfTrue(ctx *InstructionContext) {
	params := ctx.Parameters
	if *params[0] != 0 {
		ctx.NextInstPtr = *params[1]
	}
	// Otherwise it does nothing
}

// --- Day 5 Part 2 ---
// Opcode 6 is jump-if-false: if the first parameter is zero, it sets the
// instruction pointer to the value from the second parameter. Otherwise, it
// does nothing.
func I06_JumpIfFalse(ctx *InstructionContext) {
	params := ctx.Parameters
	if *params[0] == 0 {
		ctx.NextInstPtr = *params[1]
	}
	// Otherwise it does nothing
}

// --- Day 5 Part 2 ---
// Opcode 7 is less than: if the first parameter is less than the second
// parameter, it stores 1 in the position given by the third parameter.
// Otherwise, it stores 0.
func I07_LessThan(ctx *InstructionContext) {
	params := ctx.Parameters
	if *params[0] < *params[1] {
		*params[2] = 1
	} else {
		*params[2] = 0
	}
}

// --- Day 5 Part 2 ---
// Opcode 8 is equals: if the first parameter is equal to the second parameter,
// it stores 1 in the position given by the third parameter. Otherwise, it
// stores 0.
func I08_Equals(ctx *InstructionContext) {
	params := ctx.Parameters
	if *params[0] == *params[1] {
		*params[2] = 1
	} else {
		*params[2] = 0
	}
}

// --- Day 9 ---
// Opcode 9 adjusts the relative base by the value of its only parameter. The
// relative base increases (or decreases, if the value is negative) by the value
// of the parameter.
func I09_AdjustRelativeBase(ctx *InstructionContext) {
	ctx.Machine.RelativeBase += *ctx.Parameters[0]
}

// --- Day 2 ---
// Opcode 99 means that the program is finished and should immediately halt.
func I99_HCF(ctx *InstructionContext) {
	ctx.Machine.State.Set(OnFire)
	close(ctx.Machine.Output)
}