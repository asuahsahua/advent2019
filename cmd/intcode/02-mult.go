package intcode

const INST_MULT = 2

// --- Day 2 ---
// Opcode 2 works exactly like opcode 1, except it multiplies the two inputs
// instead of adding them. Again, the three integers after the opcode indicate
// where the inputs and outputs are, not their values.
func (ctx InstructionContext) I02_Mult() {
	p := ctx.Parameters
	*p[2] = *p[0] * *p[1]
}