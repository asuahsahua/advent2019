package intcode

const INST_HCF = 99

func I99_HCF(ctx *InstructionContext) {
	ctx.Machine.OnFire = true
	close(ctx.Machine.Output)
}