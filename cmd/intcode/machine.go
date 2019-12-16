package intcode

import (
	. "github.com/asuahsahua/advent2019/cmd/common"
)

const (
	CHAN_BUF = 2048
)

type IntcodeMachine struct{
	Memory  []int64 // Memory
	InstPtr int64   // Instruction Pointer
	RelativeBase int64 // Relative Base for Relative Mode parameters
	OnFire  bool  // Has caught fire?

	Input   chan int64
	Output  chan int64
}

func NewIntcodeMachine(program []int64) *IntcodeMachine {
	// Day 9: "The computer's available memory should be much larger than the initial program"
	// Just give 4KB of memory to each program for now
	memory := make([]int64, 4 * 1024)
	copy(memory, program)

	return &IntcodeMachine{
		InstPtr: 0,
		// Day 9: The relative base starts at 0
		RelativeBase: 0,
		Memory: memory,
		OnFire: false,

		// Allow a buffered width of 10 for now, for laziness. 
		// Hopefully things don't need to be buffered much more than this.
		// edit: so much for that
		Input: make(chan int64, CHAN_BUF),
		Output: make(chan int64, CHAN_BUF),
	}
}

func NewIntcodeMachineStr(program string) *IntcodeMachine {
	return NewIntcodeMachine(
		CommaSeparatedToInt64(program),
	)
}

func RunProgram(program string, input int64) (output int64) {
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