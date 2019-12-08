package intcode

import (
	"fmt"
	. "github.com/asuahsahua/advent2019/cmd/common"
)

type IntcodeMachine struct{
	PC      int   // Instruction Pointer
	Memory  []int // Memory
	OnFire  bool  // Has caught fire?
}

var instructions map[int]instruction = map[int]instruction{
	INST_ADD:  I01_Add,
	INST_MULT: I02_Mult,
	INST_HCF:  I99_HCF,
}

func NewIntcodeMachine(program []int) *IntcodeMachine {
	memory := make([]int, len(program))
	copy(memory, program)

	return &IntcodeMachine{
		PC: 0,
		Memory: memory,
		OnFire: false,
	}
}

func NewIntcodeMachineStr(program string) *IntcodeMachine {
	programBytes := CommaSeparatedToInt(program)

	return NewIntcodeMachine(programBytes)
}

func (m *IntcodeMachine) Run() {
	for m.OnFire == false {
		m.RunStep()
	}
}

func (m *IntcodeMachine) RunStep() {
	inst := instructions[m.Memory[m.PC]]
	if inst == nil {
		panic(fmt.Sprintf("Could not find instruction type! %d", m.Memory[m.PC]))
	}
	argCount := inst.ArgCount()

	switch v := inst.(type) {
	case Instruction3:
		v.code(m, m.Memory[m.PC + 1], m.Memory[m.PC + 2], m.Memory[m.PC + 3])
		break;
	case Instruction0:
		v.code(m)
		break;
	default:
		fmt.Printf("Memory: %v", m.Memory)
		panic("Unrecognized instruction type found")
	}

	m.PC += 1 + argCount
}