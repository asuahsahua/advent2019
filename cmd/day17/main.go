package main

import (
	. "github.com/asuahsahua/advent2019/cmd/common"
	"github.com/asuahsahua/advent2019/cmd/intcode"
)

// --- Day 17: Set and Forget ---

// An early warning system detects an incoming solar flare and automatically
// activates the ship's electromagnetic shield. Unfortunately, this has cut off
// the Wi-Fi for many small robots that, unaware of the impending danger, are
// now trapped on exterior scaffolding on the unsafe side of the shield. To
// rescue them, you'll have to act quickly!

// The only tools at your disposal are some wired cameras and a small vacuum
// robot currently asleep at its charging station. The video quality is poor,
// but the vacuum robot has a needlessly bright LED that makes it easy to spot
// no matter where it is.
type WiredCamera struct {
}

type VacuumRobot struct {
	Awake bool
	NeedlesslyBrightLED bool
}

var SmallVacuumRobot VacuumRobot = VacuumRobot{
	Awake: false,
	NeedlesslyBrightLED: true,
}

// An Intcode program, the Aft Scaffolding Control and Information Interface
// (ASCII, your puzzle input), provides access to the cameras and the vacuum
// robot. Currently, because the vacuum robot is asleep, you can only access the
// cameras.
var ASCII_Program string = `1,330,331,332,109,3594,1102,1182,1,15,1102,1,1487,24,1001,0,0,570,1006,570,36,1001,571,0,0,1001,570,-1,570,1001,24,1,24,1106,0,18,1008,571,0,571,1001,15,1,15,1008,15,1487,570,1006,570,14,21102,1,58,0,1106,0,786,1006,332,62,99,21101,0,333,1,21102,73,1,0,1106,0,579,1102,1,0,572,1102,0,1,573,3,574,101,1,573,573,1007,574,65,570,1005,570,151,107,67,574,570,1005,570,151,1001,574,-64,574,1002,574,-1,574,1001,572,1,572,1007,572,11,570,1006,570,165,101,1182,572,127,1001,574,0,0,3,574,101,1,573,573,1008,574,10,570,1005,570,189,1008,574,44,570,1006,570,158,1105,1,81,21101,0,340,1,1106,0,177,21102,1,477,1,1106,0,177,21101,0,514,1,21101,0,176,0,1105,1,579,99,21102,184,1,0,1106,0,579,4,574,104,10,99,1007,573,22,570,1006,570,165,1001,572,0,1182,21102,375,1,1,21102,211,1,0,1106,0,579,21101,1182,11,1,21101,222,0,0,1106,0,979,21101,0,388,1,21102,233,1,0,1105,1,579,21101,1182,22,1,21102,244,1,0,1105,1,979,21102,401,1,1,21102,255,1,0,1106,0,579,21101,1182,33,1,21102,1,266,0,1105,1,979,21102,1,414,1,21101,0,277,0,1106,0,579,3,575,1008,575,89,570,1008,575,121,575,1,575,570,575,3,574,1008,574,10,570,1006,570,291,104,10,21101,1182,0,1,21101,0,313,0,1105,1,622,1005,575,327,1101,1,0,575,21102,327,1,0,1105,1,786,4,438,99,0,1,1,6,77,97,105,110,58,10,33,10,69,120,112,101,99,116,101,100,32,102,117,110,99,116,105,111,110,32,110,97,109,101,32,98,117,116,32,103,111,116,58,32,0,12,70,117,110,99,116,105,111,110,32,65,58,10,12,70,117,110,99,116,105,111,110,32,66,58,10,12,70,117,110,99,116,105,111,110,32,67,58,10,23,67,111,110,116,105,110,117,111,117,115,32,118,105,100,101,111,32,102,101,101,100,63,10,0,37,10,69,120,112,101,99,116,101,100,32,82,44,32,76,44,32,111,114,32,100,105,115,116,97,110,99,101,32,98,117,116,32,103,111,116,58,32,36,10,69,120,112,101,99,116,101,100,32,99,111,109,109,97,32,111,114,32,110,101,119,108,105,110,101,32,98,117,116,32,103,111,116,58,32,43,10,68,101,102,105,110,105,116,105,111,110,115,32,109,97,121,32,98,101,32,97,116,32,109,111,115,116,32,50,48,32,99,104,97,114,97,99,116,101,114,115,33,10,94,62,118,60,0,1,0,-1,-1,0,1,0,0,0,0,0,0,1,12,20,0,109,4,1202,-3,1,587,20101,0,0,-1,22101,1,-3,-3,21102,1,0,-2,2208,-2,-1,570,1005,570,617,2201,-3,-2,609,4,0,21201,-2,1,-2,1106,0,597,109,-4,2106,0,0,109,5,2101,0,-4,630,20101,0,0,-2,22101,1,-4,-4,21102,1,0,-3,2208,-3,-2,570,1005,570,781,2201,-4,-3,653,20102,1,0,-1,1208,-1,-4,570,1005,570,709,1208,-1,-5,570,1005,570,734,1207,-1,0,570,1005,570,759,1206,-1,774,1001,578,562,684,1,0,576,576,1001,578,566,692,1,0,577,577,21101,0,702,0,1106,0,786,21201,-1,-1,-1,1105,1,676,1001,578,1,578,1008,578,4,570,1006,570,724,1001,578,-4,578,21102,1,731,0,1105,1,786,1106,0,774,1001,578,-1,578,1008,578,-1,570,1006,570,749,1001,578,4,578,21101,0,756,0,1106,0,786,1105,1,774,21202,-1,-11,1,22101,1182,1,1,21101,0,774,0,1106,0,622,21201,-3,1,-3,1106,0,640,109,-5,2105,1,0,109,7,1005,575,802,21001,576,0,-6,20102,1,577,-5,1105,1,814,21101,0,0,-1,21101,0,0,-5,21102,1,0,-6,20208,-6,576,-2,208,-5,577,570,22002,570,-2,-2,21202,-5,49,-3,22201,-6,-3,-3,22101,1487,-3,-3,1202,-3,1,843,1005,0,863,21202,-2,42,-4,22101,46,-4,-4,1206,-2,924,21101,1,0,-1,1106,0,924,1205,-2,873,21102,35,1,-4,1105,1,924,1202,-3,1,878,1008,0,1,570,1006,570,916,1001,374,1,374,1202,-3,1,895,1102,1,2,0,1201,-3,0,902,1001,438,0,438,2202,-6,-5,570,1,570,374,570,1,570,438,438,1001,578,558,922,20101,0,0,-4,1006,575,959,204,-4,22101,1,-6,-6,1208,-6,49,570,1006,570,814,104,10,22101,1,-5,-5,1208,-5,43,570,1006,570,810,104,10,1206,-1,974,99,1206,-1,974,1102,1,1,575,21101,973,0,0,1106,0,786,99,109,-7,2106,0,0,109,6,21102,0,1,-4,21101,0,0,-3,203,-2,22101,1,-3,-3,21208,-2,82,-1,1205,-1,1030,21208,-2,76,-1,1205,-1,1037,21207,-2,48,-1,1205,-1,1124,22107,57,-2,-1,1205,-1,1124,21201,-2,-48,-2,1105,1,1041,21102,1,-4,-2,1105,1,1041,21102,1,-5,-2,21201,-4,1,-4,21207,-4,11,-1,1206,-1,1138,2201,-5,-4,1059,1202,-2,1,0,203,-2,22101,1,-3,-3,21207,-2,48,-1,1205,-1,1107,22107,57,-2,-1,1205,-1,1107,21201,-2,-48,-2,2201,-5,-4,1090,20102,10,0,-1,22201,-2,-1,-2,2201,-5,-4,1103,2102,1,-2,0,1105,1,1060,21208,-2,10,-1,1205,-1,1162,21208,-2,44,-1,1206,-1,1131,1105,1,989,21101,439,0,1,1105,1,1150,21101,477,0,1,1105,1,1150,21101,0,514,1,21101,0,1149,0,1105,1,579,99,21102,1,1157,0,1106,0,579,204,-2,104,10,99,21207,-3,22,-1,1206,-1,1138,1202,-5,1,1176,2102,1,-4,0,109,-6,2105,1,0,14,9,40,1,7,1,40,1,7,1,40,1,7,1,40,1,7,1,40,1,7,1,40,1,7,1,40,1,7,1,40,1,3,5,1,7,5,13,14,1,3,1,5,1,5,1,5,1,11,1,14,1,3,1,5,1,5,1,5,1,11,1,14,1,3,1,5,1,5,1,5,1,11,1,14,7,3,1,3,13,7,1,18,1,1,1,3,1,3,1,1,1,5,1,3,1,7,1,18,1,1,1,3,13,3,1,7,1,18,1,1,1,7,1,1,1,9,1,7,1,18,1,1,1,7,1,1,1,9,9,18,1,1,1,7,1,1,1,32,13,1,1,1,1,3,7,22,1,3,1,1,1,5,1,1,1,1,1,3,1,5,1,8,13,1,1,3,13,3,1,5,1,8,1,13,1,5,1,5,1,1,1,5,1,5,1,8,1,13,1,5,1,5,1,1,1,5,1,5,1,8,1,13,1,5,1,5,1,1,1,5,1,5,1,8,1,13,7,5,1,1,13,8,1,25,1,7,1,14,1,5,9,11,1,7,1,14,1,5,1,7,1,11,1,7,1,14,1,5,1,7,1,11,1,7,1,14,1,5,1,7,1,11,1,7,1,14,1,5,1,7,13,7,7,8,1,5,1,33,1,8,7,33,1,48,1,48,1,48,1,48,1,48,1,48,1,48,1,48,1,48,1,40,9,8`

type ASCII struct {
	Brain *intcode.IntcodeMachine
	Cameras []WiredCamera
	VacuumRobot VacuumRobot
	Screen map[Point2D]ASCIICode
}

func NewASCII() *ASCII {
	return &ASCII{
		Brain: intcode.NewIntcodeMachineStr(ASCII_Program),
		Cameras: make([]WiredCamera, 0),
		VacuumRobot: SmallVacuumRobot, 
		Screen: make(map[Point2D]ASCIICode),
	}
}

// Running the ASCII program on your Intcode computer will provide the current
// view of the scaffolds. In the camera output, # represents a scaffold and .
// represents open space.
type ASCIICode byte
const (
	Scaffold ASCIICode = '#';
	OpenSpace ASCIICode = '.';
	NewLine ASCIICode = '\n';
)

// The vacuum robot is visible as ^, v, <, or >
const (
	RobotUp ASCIICode = '^';
	RobotDown ASCIICode = 'v';
	RobotRight ASCIICode = '>';
	RobotLeft ASCIICode = '<';
	RobotDead ASCIICode = 'X';
)

func (a *ASCII) ReadScreen() map[Point2D]ASCIICode {
	a.Screen = make(map[Point2D]ASCIICode)
	pos := Point2DI(0, 0)
	for out := range a.Brain.Output {
		chr := ASCIICode(out)
		switch chr {
		case NewLine:
			pos = Point2DI(0, pos.Y + 1)
		default:
			a.Screen[pos] = chr
			pos = pos.Add(Point2DI(1, 0))
		}
	}

	return a.Screen
}

// To calibrate the cameras, you need the sum of the alignment parameters. Run
// your ASCII program. What is the sum of the alignment parameters for the
// scaffold intersections?
func (a *ASCII) FindAlignmentParameters() []int {
	junctions := make([]int, 0)

	OUTER:
	for point, val := range a.Screen {
		if val != Scaffold {
			continue
		}

		for _, adj := range point.CardinalAdjacents() {
			if a.Screen[adj] != Scaffold {
				continue OUTER
			}
		}

		junctions = append(junctions, point.X * point.Y)
	}

	return junctions
}

func (a *ASCII) PrintScreen() {
	points := make([]Point2D, 0)
	for pt, _ := range a.Screen {
		points = append(points, pt)
	}

	bbox := ResolveBoundingBox(points)
	bbox.PrintPoints(func(pt Point2D) byte {
		return byte(a.Screen[pt])
	})
}

func main() {
	ascii := NewASCII()
	go ascii.Brain.Run()
	ascii.ReadScreen()

	ascii.PrintScreen()

	alignments := ascii.FindAlignmentParameters()
	Part1("%d", SumInts(alignments))
}