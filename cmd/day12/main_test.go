package main

import (
	"fmt"
	. "github.com/asuahsahua/advent2019/cmd/common"
	. "github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestConfirmedAnswer(t *testing.T) {
	jupiter := NewJupiterStr(PuzzleInput)
	jupiter.StepTime(1000)
	Equal(t, 7722, jupiter.TotalEnergy())

	// How many steps does it take to reach the first state that exactly matches a
	// previous state?
	jup2 := NewJupiterStr(PuzzleInput)
	Equal(t, 292653556339368, jup2.FindMoonPeriod())
}

func TestExample1(t *testing.T) {
	jupiter := NewJupiterStr(Ex1Input)
	Equal(t, ParseMoonsState(Ex1Time0), jupiter.Moons)

	jupiter.StepTime(1)
	Equal(t, ParseMoonsState(Ex1Time1), jupiter.Moons)
	jupiter.StepTime(1)
	Equal(t, ParseMoonsState(Ex1Time2), jupiter.Moons)
	jupiter.StepTime(1)
	Equal(t, ParseMoonsState(Ex1Time3), jupiter.Moons)
	jupiter.StepTime(1)
	Equal(t, ParseMoonsState(Ex1Time4), jupiter.Moons)
	jupiter.StepTime(1)
	Equal(t, ParseMoonsState(Ex1Time5), jupiter.Moons)
	jupiter.StepTime(1)
	Equal(t, ParseMoonsState(Ex1Time6), jupiter.Moons)
	jupiter.StepTime(1)
	Equal(t, ParseMoonsState(Ex1Time7), jupiter.Moons)
	jupiter.StepTime(1)
	Equal(t, ParseMoonsState(Ex1Time8), jupiter.Moons)
	jupiter.StepTime(1)
	Equal(t, ParseMoonsState(Ex1Time9), jupiter.Moons)
	jupiter.StepTime(1)
	Equal(t, ParseMoonsState(Ex1Time10), jupiter.Moons)

	Equal(t, 179, jupiter.TotalEnergy())
}

func TestExample2(t *testing.T) {
	jupiter := NewJupiterStr(Ex2Input)
	Equal(t, ParseMoonsState(Ex2Time00), jupiter.Moons)

	jupiter.StepTime(10)
	Equal(t, ParseMoonsState(Ex2Time10), jupiter.Moons)
	jupiter.StepTime(10)
	Equal(t, ParseMoonsState(Ex2Time20), jupiter.Moons)
	jupiter.StepTime(10)
	Equal(t, ParseMoonsState(Ex2Time30), jupiter.Moons)
	jupiter.StepTime(10)
	Equal(t, ParseMoonsState(Ex2Time40), jupiter.Moons)
	jupiter.StepTime(10)
	Equal(t, ParseMoonsState(Ex2Time50), jupiter.Moons)
	jupiter.StepTime(10)
	Equal(t, ParseMoonsState(Ex2Time60), jupiter.Moons)
	jupiter.StepTime(10)
	Equal(t, ParseMoonsState(Ex2Time70), jupiter.Moons)
	jupiter.StepTime(10)
	Equal(t, ParseMoonsState(Ex2Time80), jupiter.Moons)
	jupiter.StepTime(10)
	Equal(t, ParseMoonsState(Ex2Time90), jupiter.Moons)
	jupiter.StepTime(10)
	Equal(t, ParseMoonsState(Ex2Time100), jupiter.Moons)

	Equal(t, 1940, jupiter.TotalEnergy())
}

func TestPart2Example1(t *testing.T) {
	jupiter := NewJupiterStr(Ex1Input)
	Equal(t, 2772, jupiter.FindMoonPeriod())
}

func TestPart2Example2(t *testing.T) {
	jupiter := NewJupiterStr(Ex2Input)
	Equal(t, 4686774924, jupiter.FindMoonPeriod())
}

var Ex1Input = `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`

var Ex1Time0 = `pos=<x=-1, y=  0, z= 2>, vel=<x= 0, y= 0, z= 0>
pos=<x= 2, y=-10, z=-7>, vel=<x= 0, y= 0, z= 0>
pos=<x= 4, y= -8, z= 8>, vel=<x= 0, y= 0, z= 0>
pos=<x= 3, y=  5, z=-1>, vel=<x= 0, y= 0, z= 0>`
var Ex1Time1 = `pos=<x= 2, y=-1, z= 1>, vel=<x= 3, y=-1, z=-1>
pos=<x= 3, y=-7, z=-4>, vel=<x= 1, y= 3, z= 3>
pos=<x= 1, y=-7, z= 5>, vel=<x=-3, y= 1, z=-3>
pos=<x= 2, y= 2, z= 0>, vel=<x=-1, y=-3, z= 1>`
var Ex1Time2 = `pos=<x= 5, y=-3, z=-1>, vel=<x= 3, y=-2, z=-2>
pos=<x= 1, y=-2, z= 2>, vel=<x=-2, y= 5, z= 6>
pos=<x= 1, y=-4, z=-1>, vel=<x= 0, y= 3, z=-6>
pos=<x= 1, y=-4, z= 2>, vel=<x=-1, y=-6, z= 2>`
var Ex1Time3 = `pos=<x= 5, y=-6, z=-1>, vel=<x= 0, y=-3, z= 0>
pos=<x= 0, y= 0, z= 6>, vel=<x=-1, y= 2, z= 4>
pos=<x= 2, y= 1, z=-5>, vel=<x= 1, y= 5, z=-4>
pos=<x= 1, y=-8, z= 2>, vel=<x= 0, y=-4, z= 0>`
var Ex1Time4 = `pos=<x= 2, y=-8, z= 0>, vel=<x=-3, y=-2, z= 1>
pos=<x= 2, y= 1, z= 7>, vel=<x= 2, y= 1, z= 1>
pos=<x= 2, y= 3, z=-6>, vel=<x= 0, y= 2, z=-1>
pos=<x= 2, y=-9, z= 1>, vel=<x= 1, y=-1, z=-1>`
var Ex1Time5 = `pos=<x=-1, y=-9, z= 2>, vel=<x=-3, y=-1, z= 2>
pos=<x= 4, y= 1, z= 5>, vel=<x= 2, y= 0, z=-2>
pos=<x= 2, y= 2, z=-4>, vel=<x= 0, y=-1, z= 2>
pos=<x= 3, y=-7, z=-1>, vel=<x= 1, y= 2, z=-2>`
var Ex1Time6 = `pos=<x=-1, y=-7, z= 3>, vel=<x= 0, y= 2, z= 1>
pos=<x= 3, y= 0, z= 0>, vel=<x=-1, y=-1, z=-5>
pos=<x= 3, y=-2, z= 1>, vel=<x= 1, y=-4, z= 5>
pos=<x= 3, y=-4, z=-2>, vel=<x= 0, y= 3, z=-1>`
var Ex1Time7 = `pos=<x= 2, y=-2, z= 1>, vel=<x= 3, y= 5, z=-2>
pos=<x= 1, y=-4, z=-4>, vel=<x=-2, y=-4, z=-4>
pos=<x= 3, y=-7, z= 5>, vel=<x= 0, y=-5, z= 4>
pos=<x= 2, y= 0, z= 0>, vel=<x=-1, y= 4, z= 2>`
var Ex1Time8 = `pos=<x= 5, y= 2, z=-2>, vel=<x= 3, y= 4, z=-3>
pos=<x= 2, y=-7, z=-5>, vel=<x= 1, y=-3, z=-1>
pos=<x= 0, y=-9, z= 6>, vel=<x=-3, y=-2, z= 1>
pos=<x= 1, y= 1, z= 3>, vel=<x=-1, y= 1, z= 3>`
var Ex1Time9 = `pos=<x= 5, y= 3, z=-4>, vel=<x= 0, y= 1, z=-2>
pos=<x= 2, y=-9, z=-3>, vel=<x= 0, y=-2, z= 2>
pos=<x= 0, y=-8, z= 4>, vel=<x= 0, y= 1, z=-2>
pos=<x= 1, y= 1, z= 5>, vel=<x= 0, y= 0, z= 2>`
var Ex1Time10 = `pos=<x= 2, y= 1, z=-3>, vel=<x=-3, y=-2, z= 1>
pos=<x= 1, y=-8, z= 0>, vel=<x=-1, y= 1, z= 3>
pos=<x= 3, y=-6, z= 1>, vel=<x= 3, y= 2, z=-3>
pos=<x= 2, y= 0, z= 4>, vel=<x= 1, y=-1, z=-1>`

var Ex2Input = `<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>`
var Ex2Time00 = `pos=<x= -8, y=-10, z=  0>, vel=<x=  0, y=  0, z=  0>
pos=<x=  5, y=  5, z= 10>, vel=<x=  0, y=  0, z=  0>
pos=<x=  2, y= -7, z=  3>, vel=<x=  0, y=  0, z=  0>
pos=<x=  9, y= -8, z= -3>, vel=<x=  0, y=  0, z=  0>`
var Ex2Time10 = `pos=<x= -9, y=-10, z=  1>, vel=<x= -2, y= -2, z= -1>
pos=<x=  4, y= 10, z=  9>, vel=<x= -3, y=  7, z= -2>
pos=<x=  8, y=-10, z= -3>, vel=<x=  5, y= -1, z= -2>
pos=<x=  5, y=-10, z=  3>, vel=<x=  0, y= -4, z=  5>`
var Ex2Time20 = `pos=<x=-10, y=  3, z= -4>, vel=<x= -5, y=  2, z=  0>
pos=<x=  5, y=-25, z=  6>, vel=<x=  1, y=  1, z= -4>
pos=<x= 13, y=  1, z=  1>, vel=<x=  5, y= -2, z=  2>
pos=<x=  0, y=  1, z=  7>, vel=<x= -1, y= -1, z=  2>`
var Ex2Time30 = `pos=<x= 15, y= -6, z= -9>, vel=<x= -5, y=  4, z=  0>
pos=<x= -4, y=-11, z=  3>, vel=<x= -3, y=-10, z=  0>
pos=<x=  0, y= -1, z= 11>, vel=<x=  7, y=  4, z=  3>
pos=<x= -3, y= -2, z=  5>, vel=<x=  1, y=  2, z= -3>`
var Ex2Time40 = `pos=<x= 14, y=-12, z= -4>, vel=<x= 11, y=  3, z=  0>
pos=<x= -1, y= 18, z=  8>, vel=<x= -5, y=  2, z=  3>
pos=<x= -5, y=-14, z=  8>, vel=<x=  1, y= -2, z=  0>
pos=<x=  0, y=-12, z= -2>, vel=<x= -7, y= -3, z= -3>`
var Ex2Time50 = `pos=<x=-23, y=  4, z=  1>, vel=<x= -7, y= -1, z=  2>
pos=<x= 20, y=-31, z= 13>, vel=<x=  5, y=  3, z=  4>
pos=<x= -4, y=  6, z=  1>, vel=<x= -1, y=  1, z= -3>
pos=<x= 15, y=  1, z= -5>, vel=<x=  3, y= -3, z= -3>`
var Ex2Time60 = `pos=<x= 36, y=-10, z=  6>, vel=<x=  5, y=  0, z=  3>
pos=<x=-18, y= 10, z=  9>, vel=<x= -3, y= -7, z=  5>
pos=<x=  8, y=-12, z= -3>, vel=<x= -2, y=  1, z= -7>
pos=<x=-18, y= -8, z= -2>, vel=<x=  0, y=  6, z= -1>`
var Ex2Time70 = `pos=<x=-33, y= -6, z=  5>, vel=<x= -5, y= -4, z=  7>
pos=<x= 13, y= -9, z=  2>, vel=<x= -2, y= 11, z=  3>
pos=<x= 11, y= -8, z=  2>, vel=<x=  8, y= -6, z= -7>
pos=<x= 17, y=  3, z=  1>, vel=<x= -1, y= -1, z= -3>`
var Ex2Time80 = `pos=<x= 30, y= -8, z=  3>, vel=<x=  3, y=  3, z=  0>
pos=<x= -2, y= -4, z=  0>, vel=<x=  4, y=-13, z=  2>
pos=<x=-18, y= -7, z= 15>, vel=<x= -8, y=  2, z= -2>
pos=<x= -2, y= -1, z= -8>, vel=<x=  1, y=  8, z=  0>`
var Ex2Time90 = `pos=<x=-25, y= -1, z=  4>, vel=<x=  1, y= -3, z=  4>
pos=<x=  2, y= -9, z=  0>, vel=<x= -3, y= 13, z= -1>
pos=<x= 32, y= -8, z= 14>, vel=<x=  5, y= -4, z=  6>
pos=<x= -1, y= -2, z= -8>, vel=<x= -3, y= -6, z= -9>`
var Ex2Time100 = `pos=<x=  8, y=-12, z= -9>, vel=<x= -7, y=  3, z=  0>
pos=<x= 13, y= 16, z= -3>, vel=<x=  3, y=-11, z= -5>
pos=<x=-29, y=-11, z= -1>, vel=<x= -3, y=  7, z=  4>
pos=<x= 16, y=-13, z= 23>, vel=<x=  7, y=  1, z=  1>`

var numRex = `\s*(-?\d+)`
var moonstateRex = regexp.MustCompile(fmt.Sprintf(
	`pos=<x=%s, y=%s, z=%s>, vel=<x=%s, y=%s, z=%s>`,
	numRex, numRex, numRex, numRex, numRex, numRex,
))

func ParseMoonsState(str string) *Moons {
	moons := NewMoons()
	for _, match := range moonstateRex.FindAllStringSubmatch(str, -1) {
		moonPos := Point3DStrs(match[1], match[2], match[3])
		moon := NewMoon(moonPos)
		moon.Velocity = Point3DStrs(match[4], match[5], match[6])
		moons.Add(moon)
	}
	return moons
}
