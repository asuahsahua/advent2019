package main

import (
	"testing"
	. "github.com/stretchr/testify/assert"
)
func TestConfirmedAnswer(t *testing.T) {
	Equal(t, 6627023, runProgram(Input1, 12, 2))

	requiredOutput := 19690720
	noun, verb := findRequiredOutput(Input1, requiredOutput)
	Equal(t, 4019, 100 * noun + verb)
}