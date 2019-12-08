package intcode

const INST_HCF = 99

var I99_HCF Instruction0 = Instruction0{
	code: func(m *IntcodeMachine) {
		m.OnFire = true
	},
}