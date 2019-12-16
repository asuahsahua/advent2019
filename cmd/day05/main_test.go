package main

import (
	"github.com/asuahsahua/advent2019/cmd/intcode"
	"testing"
	. "github.com/stretchr/testify/assert"
)
func TestConfirmedAnswer(t *testing.T) {
	Equal(t, int64(12896948), RunDiagnostic())
	Equal(t, int64(7704130), intcode.RunProgram(Input, 5))
}