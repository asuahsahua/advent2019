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
		m.State.Set(Paused)
		m.MemoryLock.Unlock()

		m.PauseGroup.Wait()

		m.MemoryLock.Lock()
		m.State.Set(Running)
	default:
		Panic("Unhandled interrupt %d", i)
	}
}

// Send the suspend signal to the machine, block until suspended
// To resume the machine execute m.Run() again!
func (m *IntcodeMachine) Suspend() {
	m.Interrupt <- Suspend
	m.State.WaitFor(Suspended)
}

// Send the pause signal to the machine, block until paused
// To resume send the m.Resume() signal
func (m *IntcodeMachine) Pause() {
	m.PauseGroup.Add(1)
	m.Interrupt <- Pause
	m.State.WaitFor(Paused)
}

// Resume the machine from a paused state, block until running
func (m *IntcodeMachine) Resume() {
	PanicIf(m.State.Get() != Paused, "Cannot resume a machine that is not paused!")
	m.PauseGroup.Done()
	m.State.WaitFor(Running)
}