package intcode

const INST_HCF = 99

func (ctx InstructionContext) I99_HCF() {
	ctx.Machine.OnFire = true
}