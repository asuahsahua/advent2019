package main

import (
	. "github.com/asuahsahua/advent2019/cmd/common"
	"github.com/asuahsahua/advent2019/cmd/intcode"
)

// --- Day 15: Oxygen System ---

// Out here in deep space, many things can go wrong. Fortunately, many of those
// things have indicator lights. Unfortunately, one of those lights is lit: the
// oxygen system for part of the ship has failed!

// According to the readouts, the oxygen system must have failed days ago after
// a rupture in oxygen tank two; that section of the ship was automatically
// sealed once oxygen levels went dangerously low. A single remotely-operated
// repair droid is your only option for fixing the oxygen system.

// The Elves' care package included an Intcode program (your puzzle input) that
// you can use to remotely control the repair droid. By running that program,
// you can direct the repair droid to the oxygen system and fix the problem.
type DroidRemote struct {
	Program *intcode.IntcodeMachine
	Input chan Direction
	Output chan DroidStatusCode
}
var DroidRemoteCode string = `3,1033,1008,1033,1,1032,1005,1032,31,1008,1033,2,1032,1005,1032,58,1008,1033,3,1032,1005,1032,81,1008,1033,4,1032,1005,1032,104,99,102,1,1034,1039,102,1,1036,1041,1001,1035,-1,1040,1008,1038,0,1043,102,-1,1043,1032,1,1037,1032,1042,1105,1,124,101,0,1034,1039,1001,1036,0,1041,1001,1035,1,1040,1008,1038,0,1043,1,1037,1038,1042,1105,1,124,1001,1034,-1,1039,1008,1036,0,1041,101,0,1035,1040,1002,1038,1,1043,102,1,1037,1042,1106,0,124,1001,1034,1,1039,1008,1036,0,1041,1002,1035,1,1040,101,0,1038,1043,1002,1037,1,1042,1006,1039,217,1006,1040,217,1008,1039,40,1032,1005,1032,217,1008,1040,40,1032,1005,1032,217,1008,1039,35,1032,1006,1032,165,1008,1040,1,1032,1006,1032,165,1101,0,2,1044,1105,1,224,2,1041,1043,1032,1006,1032,179,1101,1,0,1044,1106,0,224,1,1041,1043,1032,1006,1032,217,1,1042,1043,1032,1001,1032,-1,1032,1002,1032,39,1032,1,1032,1039,1032,101,-1,1032,1032,101,252,1032,211,1007,0,71,1044,1105,1,224,1102,0,1,1044,1106,0,224,1006,1044,247,101,0,1039,1034,101,0,1040,1035,101,0,1041,1036,101,0,1043,1038,1001,1042,0,1037,4,1044,1105,1,0,63,79,32,16,21,23,90,91,50,57,98,31,96,21,59,30,88,68,89,15,28,86,14,75,41,29,86,4,80,51,46,48,68,93,74,17,76,18,32,36,80,2,77,80,9,98,38,82,65,93,76,29,23,89,97,13,75,35,2,91,73,86,69,90,9,78,84,6,16,98,97,91,66,41,99,56,35,78,15,85,67,77,55,96,59,20,88,24,80,48,85,79,92,23,68,67,99,98,96,57,20,32,90,20,6,79,33,97,21,58,90,41,83,83,7,64,14,8,92,59,83,13,96,95,51,89,41,72,51,82,60,34,81,56,77,10,4,14,61,74,94,87,3,86,52,84,92,35,88,28,78,17,57,72,85,67,56,82,83,54,89,33,4,84,3,66,45,85,16,22,74,94,75,57,68,80,86,94,18,27,53,90,72,38,95,34,20,99,98,40,95,93,55,46,7,29,87,32,56,21,98,30,88,95,77,24,73,95,14,85,2,66,73,30,85,8,69,78,75,93,4,76,56,51,89,99,51,94,14,72,39,85,96,98,37,37,75,79,61,73,96,4,97,41,92,68,58,76,29,29,78,97,44,73,67,75,85,18,1,2,9,99,10,98,19,11,73,67,86,1,94,35,29,16,99,27,35,76,42,60,99,43,28,74,11,74,91,81,11,13,91,97,75,80,68,51,81,81,77,51,72,75,59,85,62,83,91,9,20,83,57,61,31,94,80,26,52,93,86,87,78,39,46,74,86,55,24,87,95,16,82,49,75,11,73,92,64,69,43,82,41,50,24,98,8,3,73,77,19,49,99,29,96,35,86,82,60,65,36,92,89,84,69,58,95,31,67,84,44,78,24,80,46,48,98,39,94,10,78,89,95,28,82,41,97,88,23,83,67,42,97,44,78,83,28,29,66,94,45,61,37,79,55,79,30,95,45,47,76,18,84,81,93,29,90,90,86,13,86,18,47,86,87,70,1,92,98,16,70,21,54,85,54,29,73,76,80,59,84,92,16,81,87,33,96,86,29,18,84,42,60,94,67,59,89,26,42,91,42,75,58,95,81,82,38,49,85,52,43,93,90,41,88,85,12,37,77,78,95,35,87,35,35,55,92,72,26,76,19,96,19,87,66,97,81,85,58,58,74,39,74,43,51,90,48,77,56,78,16,81,57,34,95,72,18,6,75,16,61,89,56,59,76,35,18,98,76,5,75,11,86,93,51,94,6,76,84,26,82,10,29,95,74,20,74,78,5,63,14,96,84,54,55,75,85,24,95,72,54,49,92,78,22,95,97,58,70,87,28,41,88,25,75,7,29,95,67,32,82,80,81,41,63,69,56,10,81,75,8,18,94,56,67,18,83,56,64,93,84,60,73,95,13,72,4,96,97,40,77,35,62,78,77,35,73,56,99,40,64,60,90,82,86,52,89,17,21,87,84,19,92,81,92,84,81,67,73,9,26,87,2,11,76,31,72,61,89,11,78,83,67,1,64,97,82,12,73,99,81,68,58,77,15,14,31,91,76,58,17,83,45,54,77,40,47,82,40,72,73,95,10,96,29,77,21,92,87,11,55,93,87,84,8,89,51,24,87,38,97,92,48,99,8,49,78,42,91,78,50,87,89,46,80,83,25,11,74,22,81,39,99,53,93,61,93,65,83,80,35,2,85,27,33,95,24,99,86,23,89,9,26,75,66,81,29,75,20,89,8,97,17,73,63,82,73,90,32,92,68,82,59,93,48,78,67,98,34,91,32,82,73,74,2,77,16,90,61,75,30,92,0,0,21,21,1,10,1,0,0,0,0,0,0`
func NewDroidRemote() *DroidRemote {
	return &DroidRemote{
		Program: intcode.NewIntcodeMachineStr(DroidRemoteCode),
		Input: make(chan Direction),
		Output: make(chan DroidStatusCode),
	}
}

/* going to use snapshots for this, so add that capability to the remote */
func (r *DroidRemote) Snapshot() *DroidRemote {
	ss := NewDroidRemote()
	ss.Program = r.Program.Snapshot()
	return ss
}

// The remote control program executes the following steps in a loop forever:
func (r DroidRemote) Run() {
	go r.Program.Run()

	for {
		// Accept a movement command via an input instruction.
		movement := <- r.Input

		// Send the movement command to the repair droid.
		r.Program.Input <- int64(movement)

		// Wait for the repair droid to finish the movement operation.
		statusCode := <- r.Program.Output

		// Report on the status of the repair droid via an output instruction.
		r.Output <- DroidStatusCode(statusCode)
	}
}

// Only four movement commands are understood: north, south, west, and east
type Direction int64
const (
	North Direction = 1
	South Direction = 2
	West  Direction = 3
	East  Direction = 4
)

// going to think in cartesan coordinates
func direction(from, to Point2D) Direction {
	vec := to.Sub(from)
	switch vec {
	case Point2D{X: 1, Y: 0}:
		return East
	case Point2D{X: -1, Y: 0}:
		return West
	case Point2D{X: 0, Y: 1}:
		return North
	case Point2D{X: 0, Y: -1}:
		return South
	default:
		Panic("Unexpected vector %v", vec)
		return 0
	}
}

// The repair droid can reply with any of the following status codes
type DroidStatusCode int64
const (
	// The repair droid hit a wall. Its position has not changed.
	HitWall DroidStatusCode = 0
	// The repair droid has moved one step in the requested direction.
	Stepped DroidStatusCode = 1
	// The repair droid has moved one step in the requested direction; its new
	// position is the location of the oxygen system.
	SteppedOxygen DroidStatusCode = 2
)

// You don't know anything about the area around the repair droid, but you can
// figure it out by watching the status codes.
type Mapper struct {
	Map map[Point2D]MapperSpace
}
func NewMapper() *Mapper {
	mapping := make(map[Point2D]MapperSpace)
	mapping[Point2D{}] = MapperSpace{
		Type: Open,
		Steps: 0,
		DroidSnapshot: NewDroidRemote(),
		Position: Point2D{X: 0, Y: 0},
	}

	return &Mapper{
		Map: mapping,
	}
}

type MapperSpace struct {
	Type MapperSpaceType
	Steps int
	DroidSnapshot *DroidRemote
	Position Point2D
}

type MapperSpaceType int
const (
	Unexplored MapperSpaceType = 0
	Open MapperSpaceType = 1
	Blocked MapperSpaceType = 2
	Oxygen MapperSpaceType = 3
)

// What is the fewest number of movement commands required to move the repair
// droid from its starting position to the location of the oxygen system?
func (m *Mapper) FindOxygen() MapperSpace {
	m.FullyExplore()

	for _, space := range m.Map {
		if space.Type == Oxygen {
			return space
		}
	}

	panic("Could not find oxygen :(")
}

func (m *Mapper) FullyExplore() {
	var source Point2D

	/* a 'fully' explored space has the adjacents explored, partial has the center space explored */
	partially := []Point2D{Point2D{X:0, Y:0}}

	for len(partially) > 0 {
		/* pop the next space to explore */
		source, partially = partially[0], partially[1:]
		adjacents := source.CardinalAdjacents()

		for _, adj := range adjacents {
			if m.Explore(source, adj).Type == Open {
				partially = append(partially, adj)
			}
		}
	}
}

// Explore the spaces around the given point
func (m *Mapper) Explore(from Point2D, to Point2D) (tospace MapperSpace) {
	if m.Map[to].Type != Unexplored {
		return
	}

	fromspace := m.Map[from]
	PanicIf(fromspace.DroidSnapshot == nil, "Could not find snapshot for previously visited space! %v", from)

	tospace.Position = to
	tospace.Steps = fromspace.Steps + 1

	// create a snapshot droid and give it the direction we're going
	snapshot := fromspace.DroidSnapshot.Snapshot()
	go snapshot.Run()
	snapshot.Input <- direction(from, to)
	statusCode := <- snapshot.Output
	tospace.DroidSnapshot = snapshot

	switch statusCode {
	case HitWall:
		tospace.Type = Blocked
	case Stepped:
		tospace.Type = Open
	case SteppedOxygen:
		tospace.Type = Oxygen
	default:
		Panic("Unknown status code %d", statusCode)
	}

	m.Map[to] = tospace
	return
}

// You quickly repair the oxygen system; oxygen gradually fills the area.

// Oxygen starts in the location containing the repaired oxygen system. It takes
// one minute for oxygen to spread to all open locations that are adjacent to a
// location that already contains oxygen. Diagonal locations are not adjacent.

// Use the repair droid to get a complete map of the area. How many minutes will
// it take to fill with oxygen?
func (m *Mapper) UntilAirFilled() int {
	oxygen := m.FindOxygen()

	// create a new mapper from that position, setting the origin point to have
	// a new robot clone
	oxyFill := NewMapper()
	origin := Point2D{X: 0, Y: 0}
	oxyFill.Map[origin] = MapperSpace{
		Type: Open,
		Steps: 0,
		DroidSnapshot: oxygen.DroidSnapshot.Snapshot(),
		Position: origin,
	}

	oxyFill.FullyExplore()

	mostSteps := 0
	for _, space := range oxyFill.Map {
		if space.Type == Open {
			mostSteps = MaxI(mostSteps, space.Steps)
		}
	}

	return mostSteps
}

func main() {
	m := NewMapper()

	Part1("Steps to oxygen: %d", m.FindOxygen().Steps)
	Part2("Time for oxygen to fill: %d", m.UntilAirFilled())
}