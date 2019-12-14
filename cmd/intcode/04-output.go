package intcode

const INST_OUTPUT = 4

// --- Day 5 ---
// Opcode 4 outputs the value of its only parameter. For example, the
// instruction 4,50 would output the value at address 50.
func (ctx InstructionContext) I04_Output() {
	ctx.Machine.Output = *ctx.Parameters[0]
}