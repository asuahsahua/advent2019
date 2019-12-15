package intcode

import (
	"testing"
	. "github.com/stretchr/testify/assert"
)

func TestRunAmpProgram(t *testing.T) {
	// Here are some example programs:

    // Max thruster signal 43210 (from phase setting sequence 4,3,2,1,0):
	prog1 := `3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0`
	phase1 := []int{4, 3, 2, 1, 0}
	Equal(t, 43210, AmpProgramRun(prog1, phase1))

    // Max thruster signal 54321 (from phase setting sequence 0,1,2,3,4):
	prog2 := `3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0`
	phase2 := []int{0, 1, 2, 3, 4}
	Equal(t, 54321, AmpProgramRun(prog2, phase2))

	// Max thruster signal 65210 (from phase setting sequence 1,0,4,3,2):
	prog3 := `3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0`
	phase3 := []int{1, 0, 4, 3, 2}
	Equal(t, 65210, AmpProgramRun(prog3, phase3))
}

func TestOptimizeAmpProgram(t *testing.T) {
    // Max thruster signal 43210 (from phase setting sequence 4,3,2,1,0):
	prog1 := `3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0`
	phase1 := OptimizeAmpProgram(prog1)
	Equal(t, 43210, AmpProgramRun(prog1, phase1))

    // Max thruster signal 54321 (from phase setting sequence 0,1,2,3,4):
	prog2 := `3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0`
	phase2 := OptimizeAmpProgram(prog2)
	Equal(t, 54321, AmpProgramRun(prog2, phase2))

	// Max thruster signal 65210 (from phase setting sequence 1,0,4,3,2):
	prog3 := `3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0`
	phase3 := OptimizeAmpProgram(prog3)
	Equal(t, 65210, AmpProgramRun(prog3, phase3))
}