package main

import (
	"github.com/asuahsahua/advent2019/cmd/intcode"
	"testing"
	. "github.com/stretchr/testify/assert"
)

func TestConfirmedAnswer(t *testing.T) {
	bestAmpPhase := intcode.OptimizeAmpProgram(AmplifierControllerSoftware)
	Equal(t, int64(116680), intcode.AmpProgramRun(AmplifierControllerSoftware, bestAmpPhase))

	bestFeedbackPhase := intcode.OptimizeFeedbackProgram(AmplifierControllerSoftware)
	Equal(t, int64(89603079), intcode.FeedbackProgramRun(AmplifierControllerSoftware, bestFeedbackPhase))
}