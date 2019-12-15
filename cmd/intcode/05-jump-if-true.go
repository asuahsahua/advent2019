package intcode

const INST_JUMP_IF_TRUE = 5

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