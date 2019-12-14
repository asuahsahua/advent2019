package intcode

import (
	. "github.com/asuahsahua/advent2019/cmd/common"
)

type IntcodeMachine struct{
	PC      int   // Instruction Pointer
	Memory  []int // Memory
	OnFire  bool  // Has caught fire?
	Input   int
	Output  int
}

func NewIntcodeMachine(program []int) *IntcodeMachine {
	memory := make([]int, len(program))
	copy(memory, program)

	return &IntcodeMachine{
		PC: 0,
		Memory: memory,
		OnFire: false,
	}
}

func NewIntcodeMachineStr(program string) *IntcodeMachine {
	return NewIntcodeMachine(
		CommaSeparatedToInt(program),
	)
}

func (m *IntcodeMachine) Run() {
	for m.OnFire == false {
		m.RunStep()
	}
}

func (m *IntcodeMachine) RunStep() {
	ctx := m.BuildInstructionContext(m.PC)
	ctx.Execute()
	m.PC += 1 + len(ctx.Parameters)
}