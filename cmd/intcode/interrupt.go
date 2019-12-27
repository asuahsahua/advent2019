package intcode

import (
	. "github.com/asuahsahua/advent2019/cmd/common"
)

// Interrupts allow us to 'interrupt' code flow on the interrupt channel.
// Generally, these are checked on any 'blocking' operation, such as I/O

type Interrupt int64
const (
	Suspend Interrupt = 1
	Pause Interrupt = 2
)

func (m *IntcodeMachine) handleInterrupt(i Interrupt) {
	switch i {
	case Suspend:
		m.State.Set(Suspended)
	case Pause:
		m.PauseGroup.Wait()
	default:
		Panic("Unhandled interrupt %d", i)
	}
}

func (m *IntcodeMachine) Suspend() {
	m.Interrupt <- Suspend
}

func (m *IntcodeMachine) Pause() {
	m.PauseGroup.Add(1)
	m.Interrupt <- Pause
}

func (m *IntcodeMachine) Resume() {
	m.PauseGroup.Done()
}