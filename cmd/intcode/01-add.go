package intcode

const INST_ADD = 1

// --- Day 2 ---
// Opcode 1 adds together numbers read from two positions and stores the result
// in a third position. The three integers immediately after the opcode tell you
// these three positions - the first two indicate the positions from which you
// should read the input values, and the third indicates the position at which
// the output should be stored.
func (ctx InstructionContext) I01_Add() {
	p := ctx.Parameters
	*p[2] = *p[0] + *p[1]
}