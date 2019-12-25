package main

import (
	. "github.com/asuahsahua/advent2019/cmd/common"
	"reflect"
	"regexp"
	"strconv"
)

// --- Day 12: The N-Body Problem ---

// The space near Jupiter is not a very safe place; you need to be careful of a
// big distracting red spot, extreme radiation, and a whole lot of moons
// swirling around.

// You decide to start by tracking the four largest moons: Io, Europa, Ganymede,
// and Callisto. After a brief scan, you calculate the position of each moon
// (your puzzle input).
var inputRex = regexp.MustCompile(`<x=(-?\d+), y=(-?\d+), z=(-?\d+)>`)
var PuzzleInput = `<x=-4, y=3, z=15>
<x=-11, y=-10, z=13>
<x=2, y=2, z=18>
<x=7, y=-1, z=0>`

// (to parse the puzzle input into points)
func ParseInput(str string) []Point3D {
	// <x=-1, y=0, z=2>
	points := make([]Point3D, 0)
	for _, match := range inputRex.FindAllStringSubmatch(str, -1) {
		X, xerr := strconv.Atoi(match[1])
		PanicIfErr(xerr)
		Y, yerr := strconv.Atoi(match[2])
		PanicIfErr(yerr)
		Z, zerr := strconv.Atoi(match[3])
		PanicIfErr(zerr)

		points = append(points, Point3D{
			X: X,
			Y: Y,
			Z: Z,
		})
	}

	return points
}

// You just need to simulate their motion so you can avoid them.
// Each moon has a 3-dimensional position (x, y, and z) and a 3-dimensional
// velocity. The position of each moon is given in your scan; the x, y, and z
// velocity of each moon starts at 0.
type Moon struct {
	Position Point3D
	Velocity Point3D
}

func NewMoon(position Point3D) *Moon { // :sparkles:
	return &Moon{
		Position: position,
		Velocity: Point3D{},
	}
}

func (m *Moon) Clone() *Moon {
	return &Moon{
		Position: m.Position,
		Velocity: m.Velocity,
	}
}

// Simulate the motion of the moons in time steps.
func (jup *Jupiter) StepTime(times int) {
	// Within each time step...
	for i := 0; i < times; i++ {
		jup.UpdateVelocities()
		jup.UpdatePositions()

		// Time progresses by one step once all of the positions are updated.
		jup.Time++
	}
}

// First, update the velocity of every moon by applying gravity.
func (jup *Jupiter) UpdateVelocities() {
	// To apply gravity, consider every pair of moons.

	// On each axis (x, y, and z), the velocity of each moon changes by exactly
	// +1 or -1 to pull the moons together
	countX := jup.Moons.CountBy(func(moon *Moon) int { return moon.Position.X })
	countY := jup.Moons.CountBy(func(moon *Moon) int { return moon.Position.Y })
	countZ := jup.Moons.CountBy(func(moon *Moon) int { return moon.Position.Z })

	// Small function to return a delta-velocity from a current position component
	// and count distribution for that component
	deltaV := func(posValue int, counts map[int]int) int {
		totalChange := 0
		for otherPos, count := range counts {
			if otherPos > posValue {
				totalChange += count
			} else if otherPos < posValue {
				totalChange -= count
			} // otherwise don't change anything
		}
		return totalChange
	}

	// (for each moon...)
	for _, moon := range jup.Moons.moons {
		moon.Velocity.X += deltaV(moon.Position.X, countX)
		moon.Velocity.Y += deltaV(moon.Position.Y, countY)
		moon.Velocity.Z += deltaV(moon.Position.Z, countZ)
	}
}

// (counts the moons into buckets by callback)
func (m Moons) CountBy(by func(*Moon) int) map[int]int {
	counts := make(map[int]int, 0)
	for _, moon := range m.moons {
		counts[by(moon)]++
	}
	return counts
}

// Then, once all moons' velocities have been updated, update the position
// of every moon by applying velocity.
func (jup *Jupiter) UpdatePositions() {
	for _, moon := range jup.Moons.moons {
		moon.Position = moon.Position.Add(moon.Velocity)
	}
}

// Then, it might help to calculate the total energy in the system.
func (jup Jupiter) TotalEnergy() int {
	totalEnergy := 0
	for _, moon := range jup.Moons.moons {
		totalEnergy += moon.Energy()
	}
	return totalEnergy
}

// The total energy for a single moon is its potential energy multiplied by
// its kinetic energy.
func (m Moon) Energy() int {
	return m.PotentialEnergy() * m.KineticEnergy()
}

// Potential energy is the sum of the absolute values of its position coordinates.
func (m Moon) PotentialEnergy() int {
	return m.Position.AbsSum()
}

// Kinetic energy is the sum of the absolute values of its velocity coordinates.
func (m Moon) KineticEnergy() int {
	return m.Velocity.AbsSum()
}

type Moons struct {
	moons []*Moon
}

func NewMoons() *Moons {
	return &Moons{
		moons: make([]*Moon, 0),
	}
}
func (m *Moons) Add(moon *Moon) {
	// maybe should check for collisions
	m.moons = append(m.moons, moon)
}

func (m *Moons) Clone() *Moons {
	moons := make([]*Moon, 0)
	for _, moon := range m.moons {
		moons = append(moons, moon.Clone())
	}

	return &Moons{
		moons: moons,
	}
}

// mapp runs the given callback on each moon (`map` is reserved)
func (m *Moons) mapp(cb func(Point3D) int) ([]int, []int) {
	pvals := make([]int, 0)
	vvals := make([]int, 0)
	for _, moon := range m.moons {
		pvals = append(pvals, cb(moon.Position))
		vvals = append(vvals, cb(moon.Velocity))
	}
	return pvals, vvals
}

type Jupiter struct {
	Moons *Moons
	Time  int
}

func NewJupiter() *Jupiter {
	return &Jupiter{
		Moons: NewMoons(),
		Time:  0,
	}
}
func NewJupiterStr(input string) *Jupiter {
	j := NewJupiter()
	for _, pos := range ParseInput(input) {
		j.Moons.Add(NewMoon(pos))
	}
	return j
}

func (jup *Jupiter) Clone() *Jupiter {
	clone := NewJupiter()
	clone.Time = jup.Time
	clone.Moons = jup.Moons.Clone()
	return clone
}

// FindMoonPeriodCallback finds the number of steps that it takes for the
// callback on each moon to match up again.
// For example, if you wanted to find the period for the x-axis, the callback would be:
// func (m *Moon) { return m.Position.X }
func (jup *Jupiter) FindMoonPeriodCallback(cb func(Point3D) int) int {
	clone := jup.Clone()

	startP, _ := clone.Moons.mapp(cb)
	zerosV := make([]int, len(startP)) // need the ccomponent

	for {
		clone.StepTime(1)

		currP, currV := clone.Moons.mapp(cb)
		if reflect.DeepEqual(currP, startP) && reflect.DeepEqual(currV, zerosV) {
			return clone.Time
		} // otherwise continue
	}
}

// FindMoonPeriod finds the number of steps until each moon is in their original
// location
func (jup *Jupiter) FindMoonPeriod() int {
	// Find the period of each of the position components
	xPeriod := jup.FindMoonPeriodCallback(func(p Point3D) int { return p.X })
	yPeriod := jup.FindMoonPeriodCallback(func(p Point3D) int { return p.Y })
	zPeriod := jup.FindMoonPeriodCallback(func(p Point3D) int { return p.Z })

	return LCM(xPeriod, yPeriod, zPeriod)
}

func main() {
	// What is the total energy in the system after simulating the moons given in
	// your scan for 1000 steps?
	jupiter := NewJupiterStr(PuzzleInput)
	jupiter.StepTime(1000)
	Part1("%d", jupiter.TotalEnergy())

	// How many steps does it take to reach the first state that exactly matches a
	// previous state?
	jup2 := NewJupiterStr(PuzzleInput)
	Part2("%d", jup2.FindMoonPeriod())
}
