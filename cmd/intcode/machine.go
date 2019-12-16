package intcode

import (
	. "github.com/asuahsahua/advent2019/cmd/common"
)

type IntcodeMachine struct{
	Memory  []int // Memory
	InstPtr int   // Instruction Pointer
	RelativeBase int // Relative Base for Relative Mode parameters
	OnFire  bool  // Has caught fire?

	Input   chan int
	Output  chan int
}

func NewIntcodeMachine(program []int) *IntcodeMachine {
	// Day 9: "The computer's available memory should be much larger than the initial program"
	// Just give 4KB of memory to each program for now
	memory := make([]int, 4 * 1024)
	copy(memory, program)

	return &IntcodeMachine{
		InstPtr: 0,
		// Day 9: The relative base starts at 0
		RelativeBase: 0,
		Memory: memory,
		OnFire: false,

		// Allow a buffered width of 10 for now, for laziness. 
		// Hopefully things don't need to be buffered much more than this.
		Input: make(chan int, 10),
		Output: make(chan int, 10),
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