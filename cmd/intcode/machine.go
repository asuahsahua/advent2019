package intcode

import (
	. "github.com/asuahsahua/advent2019/cmd/common"
)

type IntcodeMachine struct{
	InstPtr int   // Instruction Pointer
	Memory  []int // Memory
	OnFire  bool  // Has caught fire?

	Input   chan int
	Output  chan int
}

func NewIntcodeMachine(program []int) *IntcodeMachine {
	memory := make([]int, len(program))
	copy(memory, program)

	return &IntcodeMachine{
		InstPtr: 0,
		Memory: memory,
		OnFire: false,

		// Allow a buffered width of 1
		Input: make(chan int, 1),
		Output: make(chan int, 1),
	}
}

func NewIntcodeMachineStr(program string) *IntcodeMachine {
	return NewIntcodeMachine(
		CommaSeparatedToInt(program),
	)
}

func RunProgram(program string, input int) (output int) {
	machine := NewIntcodeMachineStr(program)
	machine.Input <- input

	machine.Run()

	return <- machine.Output 
}

func (m *IntcodeMachine) Run() {
	for m.OnFire == false {
		m.RunStep()
	}
}

func (m *IntcodeMachine) RunStep() {
	ctx := m.BuildInstructionContext(m.InstPtr)
	ctx.Execute()
	m.InstPtr = ctx.NextInstPtr
}