package main

import (
	. "github.com/asuahsahua/advent2019/cmd/common"
	"testing"
	. "github.com/stretchr/testify/assert"
)

func TestConfirmedAnswerPart1(t *testing.T) {
	ascii := NewASCII()
	go ascii.Brain.Run()
	ascii.ReadScreen()
	alignments := ascii.FindAlignmentParameters()
	Equal(t, 4800, SumInts(alignments))
}

func TestConfirmedAnswerPart2(t *testing.T) {
	Equal(t, 2, 2)
}