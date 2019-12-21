package intcode

import (
	"sync"
)

type AmpChain struct {
	Input *chan int64
	Output *chan int64
	Chain []*IntcodeMachine
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
		go func(m *IntcodeMachine) {
			m.Run()
			wg.Done()
		}(machine)
	}

	wg.Wait()

	return <- *amp.Output
}

// Build a chain of amplifiers using the same program
func NewAmpChain(prog string, size int64) *AmpChain {
	amps := make([]*IntcodeMachine, 0)
	for i := int64(0); i < size; i++ {
		machine := NewIntcodeMachineStr(prog)
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
