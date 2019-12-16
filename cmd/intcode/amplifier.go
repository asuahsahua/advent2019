package intcode

import (
	"sync"
)

type AmpChain struct {
	Input *chan int
	Output *chan int
	Chain []*IntcodeMachine
}

func AmpProgramRun(prog string, phaseSettings []int) int {
	amp := NewAmpChain(prog, len(phaseSettings))
	amp.PushPhaseSettings(phaseSettings)
	*amp.Input <- 0

	return amp.Run()
}

func (amp *AmpChain) PushPhaseSettings(phaseSettings []int) {
	// And push the appropriate phase setting to it
	for i := 0; i < len(phaseSettings); i++ {
		amp.Chain[i].Input <- phaseSettings[i]
	}
}

func (amp *AmpChain) Run() int {
	var wg sync.WaitGroup

	for _, machine := range amp.Chain {
		wg.Add(1)
		go func(m *IntcodeMachine) {
			m.Run()
			wg.Done()
		}(machine)
	}

	wg.Wait()

	return <- *amp.Output
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

func FeedbackProgramRun(prog string, phaseSettings []int) int {
	amp := NewAmpChain(prog, len(phaseSettings))

	// Need to connect the last output to the input
	amp.Chain[0].Input = amp.Chain[len(amp.Chain)-1].Output
	amp.Input = &amp.Chain[0].Input

	amp.PushPhaseSettings(phaseSettings)
	*amp.Input <- 0

	return amp.Run()
}
