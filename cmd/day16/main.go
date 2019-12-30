package main

import (
	. "github.com/asuahsahua/advent2019/cmd/common"
)

// --- Day 16: Flawed Frequency Transmission ---

// You're 3/4ths of the way through the gas giants. Not only do roundtrip
// signals to Earth take five hours, but the signal quality is quite bad as
// well. You can clean up the signal with the Flawed Frequency Transmission
// algorithm, or FFT.

// As input, FFT takes a list of numbers. In the signal you received (your
// puzzle input), each number is a single digit: data like 15243 represents
// the sequence 1, 5, 2, 4, 3.
var PuzzleInput Input = `59712692690937920492680390886862131901538154314496197364022235676243731306353384700179627460533651346711155314756853419495734284609894966089975988246871687322567664499495407183657735571812115059436153203283165299263503632551949744441033411147947509168375383038493461562836199103303184064429083384309509676574941283043596285161244885454471652448757914444304449337194545948341288172476145567753415508006250059581738670546703862905469451368454757707996318377494042589908611965335468490525108524655606907405249860972187568380476703577532080056382150009356406585677577958020969940093556279280232948278128818920216728406595068868046480073694516140765535007`

type FFT struct{
	// FFT operates in repeated phases. In each phase, a new list is constructed
	// with the same length as the input list. This new list is also used as the
	// input for the next phase.
	Value []int
	Phase *Phase
}
func NewFFT(numbers []int) *FFT {
	return &FFT{
		Value: numbers,
		Phase: BasePattern(),
	}
}

// Each element in the new list is built by multiplying every value in the input
// list by a value in a repeating pattern and then adding up the results.
func (fft *FFT) ApplyPhase(times int) []int {
	for phase := 0; phase < times; phase++ {
		out := make([]int, len(fft.Value))
		for i := 0; i < len(fft.Value); i++ {
			phase := fft.Phase.ForPosition(i, len(fft.Value))
			for k := 0; k < len(fft.Value); k++ {
				out[i] += phase[k] * fft.Value[k]
			}
			out[i] = AbsI(out[i] % 10)
		}
		fft.Value = out
	}

	return fft.Value
}

func (fft *FFT) FirstEight() string {
	return JoinI(fft.Value[0:8], "")
}

type Phase struct {
	Base []int
	Cache map[int][]int
}

// While each element in the output array uses all of the same input array
// elements, the actual repeating pattern to use depends on which output element
// is being calculated. The base pattern is 0, 1, 0, -1. 
func BasePattern() *Phase{
	return &Phase{
		Base: []int{0, 1, 0, -1},
		Cache: make(map[int][]int),
	}
}

var phaseCache map[int][]int
func (p *Phase) ForPosition(position int, length int) []int {
	if out, ok := p.Cache[position]; ok {
		return FillRepeat(out, length)
	}

	out := make([]int, 0)
	for i := 1; i <= length; i++ {
		ip := (i / (position + 1)) % len(p.Base)
		out = append(out, p.Base[ip])
	}

	p.Cache[position] = out

	// Then, repeat each value in the pattern a number of times equal to the
	// position in the output list being considered. 
	return FillRepeat(out, length)
}

func main() {
	fft := PuzzleInput.ToFFT()
	fft.ApplyPhase(100)
	Part1("First eight: %s", fft.FirstEight())
}

type Input string
func (i Input) ToFFT() *FFT {
	return NewFFT(SplitInts(string(i), ""))
}