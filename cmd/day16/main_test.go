package main

import (
	"testing"
	. "github.com/stretchr/testify/assert"
)

func TestConfirmedAnswer(t *testing.T) {
	fft := PuzzleInput.ToFFT()
	fft.ApplyPhase(100)
	Equal(t, `67481260`, fft.FirstEight())
	Equal(t, 2, 2)
}

func TestExample(t *testing.T) {
	ex := Input(`12345678`)
	fft := ex.ToFFT()

	Equal(t, "12345678", fft.FirstEight())
	fft.ApplyPhase(1)
	Equal(t, "48226158", fft.FirstEight())
	fft.ApplyPhase(1)
	Equal(t, "34040438", fft.FirstEight())
	fft.ApplyPhase(1)
	Equal(t, "03415518", fft.FirstEight())
	fft.ApplyPhase(1)
	Equal(t, "01029498", fft.FirstEight())
}

func TestOtherExamples(t *testing.T) {
	tests := map[string]string{
		"80871224585914546619083218645595": "24176176",
		"19617804207202209144916044189917": "73745418",
		"69317163492948606335995924319873": "52432133",
	}

	for input, expected := range tests {
		fft := Input(input).ToFFT()
		fft.ApplyPhase(100)
		Equal(t, expected, fft.FirstEight())
	}
}

func TestBasePattern(t *testing.T) {
	phase := BasePattern()
	Equal(t, []int{0, 1, 0, -1}, phase.Base)
	Equal(t, []int{1, 0, -1, 0, 1, 0, -1, 0, 1, 0, -1, 0}, phase.ForPosition(0, 12))
	Equal(t, []int{0, 1, 1, 0, 0, -1, -1, 0, 0, 1, 1, 0}, phase.ForPosition(1, 12))
	Equal(t, []int{0, 0, 1, 1, 1, 0, 0, 0, -1, -1, -1, 0}, phase.ForPosition(2, 12))
}

func TestInput(t *testing.T) {
	example := Input(`12345678`)
	fft := example.ToFFT()
	Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, fft.Value)
	Equal(t, []int{0, 1, 0, -1}, fft.Phase.Base)
}