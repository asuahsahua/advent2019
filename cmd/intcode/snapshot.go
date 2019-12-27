package intcode

// Snapshot creates a full snapshot of the state of the intcode machine
func (m *IntcodeMachine) Snapshot() (ss *IntcodeMachine) {
	m.Pause()

	memory := make([]int64, len(m.Memory))
	copy(memory, m.Memory)

	ss = NewIntcodeMachine(memory)
	ss.RelativeBase = m.RelativeBase
	ss.InstPtr = m.InstPtr

	m.Resume()

	return ss
}
