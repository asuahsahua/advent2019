package intcode

const INST_LESS_THAN = 7

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