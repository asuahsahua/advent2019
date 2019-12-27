package intcode

const INST_HCF = 99

func I99_HCF(ctx *InstructionContext) {
	ctx.Machine.State.Set(OnFire)
	close(ctx.Machine.Output)
}