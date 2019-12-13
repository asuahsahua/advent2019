package main

import (
	"strconv"
	"fmt"
	"strings"
	. "github.com/asuahsahua/advent2019/cmd/common"
)

func main() {

	// --- Day 3: Crossed Wires ---

	// The gravity assist was successful, and you're well on your way to the
	// Venus refuelling station. During the rush back on Earth, the fuel
	// management system wasn't completely installed, so that's next on the
	// priority list.
	
	// Opening the front panel reveals a jumble of wires. Specifically, two
	// wires are connected to a central port and extend outward on a grid. You
	// trace the path each wire takes as it leaves the central port, one wire
	// per line of text (your puzzle input).
	
	// The wires twist and turn, but the two wires occasionally cross paths. To
	// fix the circuit, you need to find the intersection point closest to the
	// central port. Because the wires are on a grid, use the Manhattan distance
	// for this measurement. While the wires do technically cross right at the
	// central port where they both start, this point does not count, nor does a
	// wire count as crossing with itself.
	
	// For example, if the first wire's path is R8,U5,L5,D3, then starting from
	// the central port (o), it goes right 8, up 5, left 5, and finally down 3:

	// ...........
	// ...........
	// ...........
	// ....+----+.
	// ....|....|.
	// ....|....|.
	// ....|....|.
	// .........|.
	// .o-------+.
	// ...........

	// Then, if the second wire's path is U7,R6,D4,L4, it goes up 7, right 6,
	// down 4, and left 4:

	// ...........
	// .+-----+...
	// .|.....|...
	// .|..+--X-+.
	// .|..|..|.|.
	// .|.-X--+.|.
	// .|..|....|.
	// .|.......|.
	// .o-------+.
	// ...........

	// These wires cross at two locations (marked X), but the lower-left one is
	// closer to the central port: its distance is 3 + 3 = 6.

	// Here are a few more examples:

	// 	R75,D30,R83,U83,L12,D49,R71,U7,L72
	// 	U62,R66,U55,R34,D71,R55,D58,R83 = distance 159
	// 	R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
	// 	U98,R91,D20,R16,D67,R40,U7,R15,U6,R7 = distance 135

	// What is the Manhattan distance from the central port to the closest
	// intersection?

	var puzzleInput = `R1005,D32,R656,U228,L629,U59,L558,D366,L659,D504,R683,U230,R689,U489,R237,U986,L803,U288,R192,D473,L490,U934,L749,D631,L333,U848,L383,D363,L641,D499,R926,D945,L520,U311,R75,D414,L97,D338,L754,U171,R601,D215,R490,U164,R158,U499,L801,U27,L671,D552,R406,U168,R12,D321,L97,U27,R833,U503,R950,U432,L688,U977,R331,D736,R231,U301,L579,U17,R984,U399,L224,U100,L266,U184,R46,D989,L851,D739,R45,D231,R893,D372,L260,U26,L697,U423,L716,D573,L269,U867,R722,U193,R889,D322,L743,U371,L986,D835,R534,U170,R946,U271,L514,D521,L781,U390,L750,D134,L767,U599,L508,U683,L426,U433,L405,U10,L359,D527,R369,D365,L405,D812,L979,D122,L782,D460,R583,U765,R502,D2,L109,D69,L560,U76,R130,D794,R197,D113,L602,D123,L190,U246,L407,D957,L35,U41,L884,D591,R38,D911,L269,D204,R332,U632,L826,D202,L984,U153,L187,U472,R272,U232,L786,U932,L618,U104,R632,D469,L868,D451,R261,U647,L211,D781,R609,D549,L628,U963,L917,D716,L218,U71,L148,U638,R34,U133,R617,U312,L215,D41,L673,U643,R379,U486,L273,D539,L294,D598,L838,D60,L158,U817,R207,U825,L601,D786,R225,D89,L417,U481,L416,U133,R261,U405,R109,U962,R104,D676,R966,U138,L343,U14,L82,U564,R73,D361,R678,D868,L273,D879,R629,U164,R228,U949,R504,D254,L662,D726,R126,D437,R569,D23,R246,U840,R457,D429,R296,U110,L984,D106,L44,U264,R801,D350,R932,D334,L252,U714,L514,U261,R632,D926,R944,U924,R199,D181,L737,U408,R636,U57,L380,D949,R557,U28,L432,D83,R829,D865,L902,D351,R71,U704,R477,D501,L882,D75,R325,D53,L990,U460,R165,D82,R577,D788,R375,U264,L178,D193,R830,D343,L394
L1003,U125,L229,U421,R863,D640,L239,U580,R342,U341,R989,U732,R51,U140,L179,U60,R483,D575,R49,U220,L284,U336,L905,U540,L392,U581,L570,U446,L817,U694,R923,U779,R624,D387,R495,D124,R862,D173,R425,D301,L550,D605,R963,U503,R571,U953,L878,D198,L256,D77,R409,D752,R921,D196,R977,U86,L842,U155,R987,D39,L224,U433,L829,D99,R558,U736,R645,D335,L52,D998,L613,D239,R470,U79,R839,D71,L753,U127,R135,D429,R729,U71,L151,U875,R668,D220,L501,D822,R306,D557,R461,U942,R59,U14,R353,D546,R409,D261,R204,U873,L847,U936,R611,U487,R474,U406,R818,U838,L301,D684,R861,D738,L265,D214,R272,D702,L145,U872,R345,D623,R200,D186,R407,U988,L608,U533,L185,D287,L549,U498,L630,U295,L425,U517,L263,D27,R697,U177,L615,U960,L553,U974,L856,U716,R126,D819,L329,D233,L212,U232,L164,D712,R316,D682,L641,U676,L535,U783,R39,U953,R39,U511,R837,U325,R391,U401,L642,U435,R626,U801,R876,D849,R448,D8,R74,U238,L186,D558,L648,D258,R262,U7,L510,U178,L183,U415,L631,D162,L521,D910,R462,U789,R885,D822,R908,D879,R614,D119,L570,U831,R993,U603,L118,U764,L414,U39,R14,U189,L415,D744,R897,U714,R326,U348,R822,U98,L357,D478,L464,D851,L545,D241,L672,U197,R156,D916,L246,U578,R4,U195,R82,D402,R327,D429,R119,U661,L184,D122,R891,D499,L808,U519,L36,U323,L259,U479,L647,D354,R891,D320,R653,U772,L158,U608,R149,U564,L164,D998,L485,U107,L145,U834,R846,D462,L391,D661,R841,U742,L597,D937,L92,U877,L350,D130,R684,U914,R400,D910,L739,U789,L188,U256,R10,U258,L965,U942,R234,D106,R852,U108,R732,U339,L955,U271,L340,U23,R373,D100,R137,U648,L130`

	wireA, wireB := InputToWires(puzzleInput)
	Part1("%d", ClosestIntersectionDistance(wireA, wireB))

	// --- Part Two ---
	// 
	// It turns out that this circuit is very timing-sensitive; you actually
	// need to minimize the signal delay.
	// 
	// To do this, calculate the number of steps each wire takes to reach each
	// intersection; choose the intersection where the sum of both wires' steps
	// is lowest. If a wire visits a position on the grid multiple times, use
	// the steps value from the first time it visits that position when
	// calculating the total value of a specific intersection.
	// 
	// The number of steps a wire takes is the total number of grid squares the
	// wire has entered to get to that location, including the intersection
	// being considered. Again consider the example from above:
	// 
	// ...........
	// .+-----+...
	// .|.....|...
	// .|..+--X-+.
	// .|..|..|.|.
	// .|.-X--+.|.
	// .|..|....|.
	// .|.......|.
	// .o-------+.
	// ...........
	// 
	// In the above example, the intersection closest to the central port is
	// reached after 8+5+5+2 = 20 steps by the first wire and 7+6+4+3 = 20 steps
	// by the second wire for a total of 20+20 = 40 steps.
	// 
	// However, the top-right intersection is better: the first wire takes only
	// 8+5+2 = 15 and the second wire takes only 7+6+2 = 15, a total of 15+15 =
	// 30 steps.
	// 
	// What is the fewest combined steps the wires must take to reach an intersection?

	Part2("%d", ClosestIntersectionBySteps(wireA, wireB))
}

// Find the distance from origin to the closest intersection between the two wires
func ClosestIntersectionDistance(a *Wire, b *Wire) int {
	// Just going to turn this into a n^2 search space because whatevs, its day 3
	nearest := 999999
	for _, intersection := range(Intersections(a, b)) {
		distance := intersection.Intersection.Manhattan(Point2D{0, 0})
		if distance < nearest {
			nearest = distance
		}
	}

	return nearest
}

// Find the closest intersection by steps, and return the number of steps to that intersection
func ClosestIntersectionBySteps(a *Wire, b *Wire) int {
	// Just going to turn this into a n^2 search space because whatevs, its day 3
	nearest := 999999
	for _, intersection := range(Intersections(a, b)) {
		fmt.Printf("%v\n", intersection)
		if intersection.CombinedSteps < nearest {
			nearest = intersection.CombinedSteps
		}
	}

	return nearest
}

// ------------ Wires -----------------
type Point2D struct{
	X int
	Y int
}

func (p1 Point2D) Manhattan(p2 Point2D) int {
	return AbsI(p1.X - p2.X) + AbsI(p1.Y - p2.Y)
}

type WireSegment struct{
	Start Point2D
	End Point2D
}

// This really only works because the line segments go in only four directions
func (ws *WireSegment) Length() int {
	return ws.Start.Manhattan(ws.End)
}

type Wire struct{
	Segments []WireSegment
}

func (w *Wire) AddSegment(seg WireSegment) {
	w.Segments = append(w.Segments, seg)
}

// ------------ Intersections --------------
// An intersection with steps counted on reaching it
type SteppedIntersection struct{
	CombinedSteps int
	Intersection Point2D
}

func Intersections(a *Wire, b *Wire) (intersections []SteppedIntersection) {
	distA := 0
	for _, segmentA := range(a.Segments) {
		distB := 0
		for _, segmentB := range(b.Segments) {
			if intersection, moreA, moreB := Intersect(segmentA, segmentB); intersection != nil {
				intersections = append(intersections, SteppedIntersection{
					CombinedSteps: distA + distB + moreA + moreB,
					Intersection: *intersection,
				})
			}
			distB += segmentB.Length()
		}
		distA += segmentA.Length()
	}
	return
}

func Intersect(wireA WireSegment, wireB WireSegment) (intersection *Point2D, distanceA int, distanceB int) {
	// lines are ax+by=c

	a1 := wireA.End.Y - wireA.Start.Y
	b1 := wireA.Start.X - wireA.End.X;
	c1 := a1 * wireA.Start.X + b1 * wireA.Start.Y

	a2 := wireB.End.Y - wireB.Start.Y
	b2 := wireB.Start.X - wireB.End.X
	c2 := a2 * wireB.Start.X + b2 * wireB.Start.Y

	det :=  a1 * b2 - a2 * b1

	if det == 0 {
		// parallel~
		return nil, 0, 0
	}
	
	// If they were both infinite lines, this is where it would intersect
	point := Point2D{
		X: (b2 * c1 - b1 * c2) / det,
		Y: (a1 * c2 - a2 * c1) / det,
	}

	// Check that the point is bounded by our segment bounding rectange
	if !IsOnWire(point, wireA) || !IsOnWire(point, wireB) {
		return nil, 0, 0
	}

	// TEMPFIX: Don't count origin
	if point.X == 0 && point.Y == 0 {
		return nil, 0, 0
	}

	return &point, 
		AbsI(point.X - wireA.Start.X) + AbsI(point.Y - wireA.Start.Y), 
		AbsI(point.X - wireB.Start.X) + AbsI(point.Y - wireB.Start.Y)
}

// Given that the point is on the 'line', check that it is bounded to the 'segment'
func IsOnWire(point Point2D, wire WireSegment) bool {
	start, end := wire.Start, wire.End
	return MinI(start.X, end.X) <= point.X && point.X <= MaxI(start.X, end.X) &&
		   MinI(start.Y, end.Y) <= point.Y && point.Y <= MaxI(start.Y, end.Y)
}

// ------------ String Parsing ------------
func InputToWires(text string) (*Wire, *Wire) {
	// Split the text on newline
	substrs := strings.Split(text, "\n")
	if len(substrs) != 2 {
		panic("Expected two lines of input")
	}

	return InputToWire(substrs[0]), InputToWire(substrs[1])
}

func InputToWire(text string) *Wire {
	split := strings.Split(text, ",")

	wire := &Wire{make([]WireSegment, 0)}
	point := Point2D{0, 0}

	for _, str := range(split) {
		direction := str[0]
		scalar, err := strconv.Atoi(str[1:])
		PanicIfErr(err)

		x, y := point.X, point.Y

		switch direction {
		case 'U':
			y += scalar
		case 'D':
			y -= scalar
		case 'R':
			x += scalar
		case 'L':
			x -= scalar
		default: panic(fmt.Sprintf("Unknown wire direction: %b", direction))
		}

		nextPoint := Point2D{x, y}

		wire.AddSegment(WireSegment{point, nextPoint})

		point = nextPoint
	}

	return wire
}
