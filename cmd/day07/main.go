package main

import (
	"sync"
	. "github.com/asuahsahua/advent2019/cmd/common"
	"github.com/asuahsahua/advent2019/cmd/intcode"
)

func main() {
	// --- Day 7: Amplification Circuit ---
	// There are five amplifiers connected in series; each one receives an input
	// signal and produces an output signal. They are connected such that the
	// first amplifier's output leads to the second amplifier's input, the
	// second amplifier's output leads to the third amplifier's input, and so
	// on. The first amplifier's input value is 0, and the last amplifier's
	// output leads to your ship's thrusters.

	//     O-------O  O-------O  O-------O  O-------O  O-------O
	// 0 ->| Amp A |->| Amp B |->| Amp C |->| Amp D |->| Amp E |-> (to thrusters)
	//     O-------O  O-------O  O-------O  O-------O  O-------O

	// The Elves have sent you some Amplifier Controller Software (your puzzle
	// input), a program that should run on your existing Intcode computer. Each
	// amplifier will need to run a copy of the program.

	// Try every combination of phase settings on the amplifiers. What is the
	// highest signal that can be sent to the thrusters?
	bestAmpPhase := OptimizeAmpProgram(AmplifierControllerSoftware)
	Part1("%d", AmpProgramRun(AmplifierControllerSoftware, bestAmpPhase))

	// It's no good - in this configuration, the amplifiers can't generate a
	// large enough output signal to produce the thrust you'll need. The Elves
	// quickly talk you through rewiring the amplifiers into a feedback loop:

	//       O-------O  O-------O  O-------O  O-------O  O-------O
	// 0 -+->| Amp A |->| Amp B |->| Amp C |->| Amp D |->| Amp E |-.
	//    |  O-------O  O-------O  O-------O  O-------O  O-------O |
	//    |                                                        |
	//    '--------------------------------------------------------+
	//                                                             |
	//                                                             v
	//                                                      (to thrusters)

	bestFeedbackPhase := OptimizeFeedbackProgram(AmplifierControllerSoftware)
	Part2("%d", FeedbackProgramRun(AmplifierControllerSoftware, bestFeedbackPhase))
}

type AmpChain struct {
	Input *chan int64
	Output *chan int64
	Chain []*intcode.IntcodeMachine
}

func AmpProgramRun(prog string, phaseSettings []int64) int64 {
	amp := NewAmpChain(prog, int64(len(phaseSettings)))
	amp.PushPhaseSettings(phaseSettings)
	*amp.Input <- 0

	return amp.Run()
}

func (amp *AmpChain) PushPhaseSettings(phaseSettings []int64) {
	// And push the appropriate phase setting to it
	for i := 0; i < len(phaseSettings); i++ {
		amp.Chain[i].Input <- phaseSettings[i]
	}
}

func (amp *AmpChain) Run() int64 {
	var wg sync.WaitGroup

	for _, machine := range amp.Chain {
		wg.Add(1)
		go func(m *intcode.IntcodeMachine) {
			m.Run()
			wg.Done()
		}(machine)
	}

	wg.Wait()

	return <- *amp.Output
}

// Build a chain of amplifiers using the same program
func NewAmpChain(prog string, size int64) *AmpChain {
	amps := make([]*intcode.IntcodeMachine, 0)
	for i := int64(0); i < size; i++ {
		machine := intcode.NewIntcodeMachineStr(prog)
		amps = append(amps, machine)

		// Connect the output of the previous with the input of the new
		if i > 0 {
			go machine.ReadFrom(amps[i - 1])
		}
	}

	return &AmpChain{
		Input:  &(amps[0].Input),
		Output: &(amps[size-1].Output),
		Chain:  amps,
	}
}

func FeedbackProgramRun(prog string, phaseSettings []int64) int64 {
	amp := NewAmpChain(prog, int64(len(phaseSettings)))

	// Need to connect the last output to the input
	amp.Chain[0].Input = amp.Chain[len(amp.Chain)-1].Output
	amp.Input = &amp.Chain[0].Input

	amp.PushPhaseSettings(phaseSettings)
	*amp.Input <- 0

	return amp.Run()
}

func OptimizeAmpProgram(prog string) []int64 {
	bestScore := int64(0)
	bestSlice := make([]int64, 5)

	I64Permutations([]int64{0, 1, 2, 3, 4}, func(perm []int64) {
		result := AmpProgramRun(prog, perm)
		if result > bestScore {
			bestScore = result
			copy(bestSlice, perm)
		}
	})

	return bestSlice
}

func OptimizeFeedbackProgram(prog string) []int64 {
	bestScore := int64(0)
	bestSlice := make([]int64, 5)

	I64Permutations([]int64{5, 6, 7, 8, 9}, func(perm []int64) {
		result := FeedbackProgramRun(prog, perm)
		if result > bestScore {
			bestScore = result
			copy(bestSlice, perm)
		}
	})

	return bestSlice
}

// The input
var AmplifierControllerSoftware string = `3,8,1001,8,10,8,105,1,0,0,21,30,47,60,81,102,183,264,345,426,99999,3,9,1002,9,5,9,4,9,99,3,9,1002,9,5,9,1001,9,4,9,1002,9,4,9,4,9,99,3,9,101,2,9,9,1002,9,4,9,4,9,99,3,9,1001,9,3,9,1002,9,2,9,101,5,9,9,1002,9,2,9,4,9,99,3,9,102,4,9,9,101,4,9,9,1002,9,3,9,101,2,9,9,4,9,99,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,1,9,4,9,99,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,101,2,9,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,99,3,9,102,2,9,9,4,9,3,9,1002,9,2,9,4,9,3,9,101,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,102,2,9,9,4,9,3,9,101,2,9,9,4,9,3,9,1002,9,2,9,4,9,99,3,9,1002,9,2,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,1,9,4,9,3,9,102,2,9,9,4,9,3,9,102,2,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,101,1,9,9,4,9,3,9,1001,9,2,9,4,9,3,9,1001,9,1,9,4,9,3,9,101,1,9,9,4,9,99`