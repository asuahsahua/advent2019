package intcode

// Run the program to completion and read all the output out
func (m *IntcodeMachine) ReadAllOutput() []int64 {
	m.Run()
	out := make([]int64, 0)

	for {
		select {
		case v := <- m.Output:
			out = append(out, v)
		default:
			return out
		}
	}
}