package intcode

const INST_EQUALS = 8

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