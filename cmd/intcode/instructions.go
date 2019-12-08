package intcode

type instruction interface{
	ArgCount() int
}

type Instruction0 struct {
	code func(m *IntcodeMachine)
}

func (i Instruction0) ArgCount() int {
	return 0
}

type Instruction3 struct {
	code func(m *IntcodeMachine, ptr1, ptr2, ptr3 int)
}

func (i Instruction3) ArgCount() int {
	return 3
}