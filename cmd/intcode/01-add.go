package intcode

const INST_ADD = 1

var I01_Add Instruction3 = Instruction3{
	code: func(m *IntcodeMachine, srcA, srcB, dest0 int) {
		m.Memory[dest0] = m.Memory[srcA] + m.Memory[srcB]
	},
}