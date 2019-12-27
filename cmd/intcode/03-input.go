package intcode

const INST_INPUT = 3

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