package main

import (
	"testing"
	. "github.com/stretchr/testify/assert"
)

func TestAmpProgramRun(t *testing.T) {
	// Here are some example programs:

    // Max thruster signal 43210 (from phase setting sequence 4,3,2,1,0):
	prog1 := `3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0`
	phase1 := []int64{4, 3, 2, 1, 0}
	Equal(t, int64(43210), AmpProgramRun(prog1, phase1))

    // Max thruster signal 54321 (from phase setting sequence 0,1,2,3,4):
	prog2 := `3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0`
	phase2 := []int64{0, 1, 2, 3, 4}
	Equal(t, int64(54321), AmpProgramRun(prog2, phase2))

	// Max thruster signal 65210 (from phase setting sequence 1,0,4,3,2):
	prog3 := `3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0`
	phase3 := []int64{1, 0, 4, 3, 2}
	Equal(t, int64(65210), AmpProgramRun(prog3, phase3))
}

func TestFeedbackProgramRun(t *testing.T) {
    // Max thruster signal 139629729 (from phase setting sequence 9,8,7,6,5):
	prog1 := `3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5`
	phase1 := []int64{9,8,7,6,5}
	Equal(t, int64(139629729), FeedbackProgramRun(prog1, phase1))

	// Max thruster signal 18216 (from phase setting sequence 9,7,8,5,6):
	prog2 := `3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10`
	phase2 := []int64{9,7,8,5,6}
	Equal(t, int64(18216), FeedbackProgramRun(prog2, phase2))
}

func TestOptimizeAmpProgram(t *testing.T) {
    // Max thruster signal 43210 (from phase setting sequence 4,3,2,1,0):
	prog1 := `3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0`
	phase1 := OptimizeAmpProgram(prog1)
	Equal(t, int64(43210), AmpProgramRun(prog1, phase1))

    // Max thruster signal 54321 (from phase setting sequence 0,1,2,3,4):
	prog2 := `3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0`
	phase2 := OptimizeAmpProgram(prog2)
	Equal(t, int64(54321), AmpProgramRun(prog2, phase2))

	// Max thruster signal 65210 (from phase setting sequence 1,0,4,3,2):
	prog3 := `3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0`
	phase3 := OptimizeAmpProgram(prog3)
	Equal(t, int64(65210), AmpProgramRun(prog3, phase3))
}

func TestOptimizeFeedbackProgram(t *testing.T) {
    // Max thruster signal 139629729 (from phase setting sequence 9,8,7,6,5):
	prog1 := `3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5`
	Equal(t, []int64{9,8,7,6,5}, OptimizeFeedbackProgram(prog1))

	// Max thruster signal 18216 (from phase setting sequence 9,7,8,5,6):
	prog2 := `3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10`
	Equal(t, []int64{9,7,8,5,6}, OptimizeFeedbackProgram(prog2))
}

func TestConfirmedAnswer(t *testing.T) {
	bestAmpPhase := OptimizeAmpProgram(AmplifierControllerSoftware)
	Equal(t, int64(116680), AmpProgramRun(AmplifierControllerSoftware, bestAmpPhase))

	bestFeedbackPhase := OptimizeFeedbackProgram(AmplifierControllerSoftware)
	Equal(t, int64(89603079), FeedbackProgramRun(AmplifierControllerSoftware, bestFeedbackPhase))
}