package intcode

const INST_JUMP_IF_FALSE = 6

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