package main

import (
	"testing"
	. "github.com/stretchr/testify/assert"
)

func TestClosestIntersectionDistance(t *testing.T) {
	input1 := `R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`
	wire1a, wire1b := InputToWires(input1)
	Equal(t, 159, ClosestIntersectionDistance(wire1a, wire1b))

	input2 := `R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`
	wire2a, wire2b := InputToWires(input2)
	Equal(t, 135, ClosestIntersectionDistance(wire2a, wire2b))

	input3 := `R8,U5,L5,D3
U7,R6,D4,L4`
	wire3a, wire3b := InputToWires(input3)
	Equal(t, 6, ClosestIntersectionDistance(wire3a, wire3b))
}

func TestClosestIntersectionBySteps(t *testing.T) {
	input1 := `R75,D30,R83,U83,L12,D49,R71,U7,L72
U62,R66,U55,R34,D71,R55,D58,R83`
	wire1a, wire1b := InputToWires(input1)
	Equal(t, 610, ClosestIntersectionBySteps(wire1a, wire1b))

	input2 := `R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
U98,R91,D20,R16,D67,R40,U7,R15,U6,R7`
	wire2a, wire2b := InputToWires(input2)
	Equal(t, 410, ClosestIntersectionBySteps(wire2a, wire2b))

	input3 := `R8,U5,L5,D3
U7,R6,D4,L4`
	wire3a, wire3b := InputToWires(input3)
	Equal(t, 30, ClosestIntersectionBySteps(wire3a, wire3b))
}

func TestConfirmedAnswer(t *testing.T) {
	wireA, wireB := InputToWires(Input)
	Equal(t, 375, ClosestIntersectionDistance(wireA, wireB))
	Equal(t, 14746, ClosestIntersectionBySteps(wireA, wireB))
}