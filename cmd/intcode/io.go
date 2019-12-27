package intcode

// Run the program to completion and read all the output out
func (m *IntcodeMachine) ReadAllOutput() []int64 {
	go m.Run()
	out := make([]int64, 0)

	for {
		v, open := <- m.Output
		if !open {
			return out
		}
		out = append(out, v)
	}
}

// (synchronously) Reads output from m2 and feeds to the input of m
// Will return when the output channel of m2 is closed
func (m *IntcodeMachine) ReadFrom(m2 *IntcodeMachine) {
	for {
		output, open := <- m2.Output
		if !open {
			return
		}

		m.Input <- output
	}
}