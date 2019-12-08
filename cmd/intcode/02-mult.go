package intcode

const INST_MULT = 2

var I02_Mult Instruction3 = Instruction3{
	code: func(m *IntcodeMachine, srcA, srcB, dest0 int) {
		m.Memory[dest0] = m.Memory[srcA] * m.Memory[srcB]
	},
}