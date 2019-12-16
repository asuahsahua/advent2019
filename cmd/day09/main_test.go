package main

import (
	"github.com/asuahsahua/advent2019/cmd/intcode"
	"testing"
	. "github.com/stretchr/testify/assert"
)

func TestConfirmedAnswer(t *testing.T) {
	// part 1
	boostTest := intcode.NewIntcodeMachineStr(BOOST)
	boostTest.Input <- 1
	testOutput := boostTest.ReadAllOutput()
	Equal(t, int64(3989758265), testOutput[len(testOutput) - 1])

	// part 2
	boost := intcode.NewIntcodeMachineStr(BOOST)
	boost.Input <- 2
	boost.Run()
	Equal(t, int64(76791), <- boost.Output)
}