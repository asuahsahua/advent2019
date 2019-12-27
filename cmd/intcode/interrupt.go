package intcode

// Interrupts allow us to 'interrupt' code flow on the interrupt channel.
// Generally, these are checked on any 'blocking' operation, such as I/O

type Interrupt int64
const (
	Suspend Interrupt = 1
)

func (m *IntcodeMachine) handleInterrupt(i Interrupt) {
	switch i {
	case Suspend:
		m.State.Set(Suspended)
	}
}

func (m *IntcodeMachine) Suspend() {
	m.Interrupt <- Suspend
}