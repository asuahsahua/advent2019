package intcode

const INST_ADJ_REL_BASE = 9

// --- Day 09 ---
// Opcode 9 adjusts the relative base by the value of its only parameter. The
// relative base increases (or decreases, if the value is negative) by the value
// of the parameter.
func I09_AdjustRelativeBase(ctx *InstructionContext) {
	ctx.Machine.RelativeBase += *ctx.Parameters[0]
}