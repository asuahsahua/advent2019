package intcode

import (
	. "github.com/asuahsahua/advent2019/cmd/common"
)

func OptimizeAmpProgram(prog string) []int {
	bestScore := 0
	bestSlice := make([]int, 5)

	IPermutations([]int{0, 1, 2, 3, 4}, func(perm []int) {
		result := AmpProgramRun(prog, perm)
		if result > bestScore {
			bestScore = result
			copy(bestSlice, perm)
		}
	})

	return bestSlice
}

type AmpChain struct {
	Input *chan int
	Output *chan int
	Chain []*IntcodeMachine
}

func AmpProgramRun(prog string, phaseSettings []int) int {
	amp := NewAmpChain(prog, len(phaseSettings))
	amp.PushPhaseSettings(phaseSettings)
	*amp.Input <- 0

	for _, machine := range amp.Chain {
		go machine.Run()
	}

	return <- *amp.Output
}

func (amp *AmpChain) PushPhaseSettings(phaseSettings []int) {
	// And push the appropriate phase setting to it
	for i := 0; i < len(phaseSettings); i++ {
		amp.Chain[i].Input <- phaseSettings[i]
	}
}

// Build a chain of amplifiers using the same program
func NewAmpChain(prog string, size int) *AmpChain {
	amps := make([]*IntcodeMachine, 0)
	for i := 0; i < size; i++ {
		machine := NewIntcodeMachineStr(prog)
		amps = append(amps, machine)

		// Connect the output of the previous with the input of the new
		if i > 0 {
			machine.Input = amps[i-1].Output
		}
	}

	return &AmpChain{
		Input:  &(amps[0].Input),
		Output: &(amps[size-1].Output),
		Chain:  amps,
	}
}