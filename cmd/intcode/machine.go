package intcode

import (
	. "github.com/asuahsahua/advent2019/cmd/common"
	"sync"
)

const (
	CHAN_BUF = 2048
)

type IntcodeMachine struct {
	// Memory Management
	Memory  []int64
	InstPtr int64
	// - RelativeBase is for relative-mode parameters (Day 9)
	RelativeBase int64

	// I/O
	Input  chan int64
	Output chan int64

	// Interrupts and state management
	Interrupt chan Interrupt
	State *MachineState
	PauseGroup *sync.WaitGroup
}

func NewIntcodeMachine(program []int64) *IntcodeMachine {
	// Day 9: "The computer's available memory should be much larger than the initial program"
	// Just give 4KB of memory to each program for now
	memory := make([]int64, 4*1024)
	copy(memory, program)

	return &IntcodeMachine{
		InstPtr: 0,
		Memory:  memory,
		// Day 9: The relative base starts at 0
		RelativeBase: 0,

		Input:  make(chan int64, CHAN_BUF),
		Output: make(chan int64, CHAN_BUF),

		Interrupt: make(chan Interrupt, CHAN_BUF),
		State: NewMachineState(),
		PauseGroup: &sync.WaitGroup{},
	}
}

func NewIntcodeMachineStr(program string) *IntcodeMachine {
	return NewIntcodeMachine(
		CommaSeparatedToInt64(program),
	)
}

// Runs the program with the given input and returns the given output
// Only really useful for dummy one-in-one-out programs, which is not going to be most programs.
func RunProgram(program string, input int64) (output int64) {
	machine := NewIntcodeMachineStr(program)
	go machine.Run()

	machine.Input <- input
	return <-machine.Output
}

func (m *IntcodeMachine) Run() {
	PanicIf(m.State.Get() != Suspended, "Cannot only run intcode machine from a suspended state")

	m.State.Set(Running)

	for m.State.Get() == Running {
		m.RunStep()
	}
}

func (m *IntcodeMachine) RunStep() {
	ctx := m.BuildInstructionContext(m.InstPtr)
	ctx.Execute()
	m.InstPtr = ctx.NextInstPtr
}
