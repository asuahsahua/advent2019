package main

import (
	"sort"
	.  "github.com/asuahsahua/advent2019/cmd/common"
)

func main() {
	// --- Day 10: Monitoring Station ---

	// You fly into the asteroid belt and reach the Ceres monitoring station.
	// The Elves here have an emergency: they're having trouble tracking all of
	// the asteroids and can't be sure they're safe.
	
	// The Elves would like to build a new monitoring station in a nearby area
	// of space; they hand you a map of all of the asteroids in that region
	// (your puzzle input).
	
	// Your job is to figure out which asteroid would be the best place to build
	// a new monitoring station. The best location is the asteroid that can
	// detect the largest number of other asteroids.

	// Find the best location for a new monitoring station. How many other
	// asteroids can be detected from that location?
	Part1("%d", BestAsteroidScoreStr(Input))

	// --- Part Two ---

	// Once you give them the coordinates, the Elves quickly deploy an Instant
	// Monitoring Station to the location and discover the worst: there are
	// simply too many asteroids.

	// The only solution is complete vaporization by giant laser.

	// Fortunately, in addition to an asteroid scanner, the new monitoring
	// station also comes equipped with a giant rotating laser perfect for
	// vaporizing asteroids. The laser starts by pointing up and always rotates
	// clockwise, vaporizing any asteroid it hits.

	// The Elves are placing bets on which will be the 200th asteroid to be
	// vaporized. Win the bet by determining which asteroid that will be; what
	// do you get if you multiply its X coordinate by 100 and then add its Y
	// coordinate? (For example, 8,2 becomes 802.)
}

type AsteroidField struct {
	Field [][]bool
}

func BestAsteroidScoreStr(fieldStr string) int {
	field := NewAsteroidFieldStr(fieldStr)
	return field.BestAsteroidScore()
}

func NewAsteroidFieldStr(str string) *AsteroidField {
	field := make([][]bool, 0)

	for _, line := range SplitLines(str) {
		row := make([]bool, 0)
		for _, c := range line {
			switch c {
			case '#':
				row = append(row, true)
			case '.':
				row = append(row, false)
			default:
				Panic("Unexpected asteroid? %c", c)
			}
		}
		field = append(field, row)
	}
	
	return &AsteroidField{
		Field: field,
	}
}

func (f *AsteroidField) EachAsteroid(callback func(point Point2D)) {
	for y, row := range f.Field {
		for x, space := range row {
			if space == true {
				callback(Point2D{x, y})
			}
		}
	}
}

func (f *AsteroidField) IsAsteroid(point Point2D) bool {
	PanicIf(len(f.Field) < point.Y, "Unexpected Y value %v", point)
	PanicIf(len(f.Field[point.Y]) < point.X, "Unexpected X value %v", point)

	return f.Field[point.Y][point.X]
}

func (f *AsteroidField) BestAsteroidScore() int {
	bestScore := 0
	// For each asteroid in our field...
	f.EachAsteroid(func(asteroid Point2D) {
		// Get the asteroids visible from that...
		visible := f.VisibleAsteroids(asteroid)
		score := len(visible)

		if score > bestScore {
			bestScore = score
		}
	})
	return bestScore
}

func (f *AsteroidField) VisibleAsteroids(asteroid Point2D) []Point2D {
	asteroids := make([]Point2D, 0)
	// Evaluate every other asteroid...
	f.EachAsteroid(func(asteroid2 Point2D) {
		if asteroid == asteroid2 {
			// Every _other_ asteroid
			return
		}

		// Can we see it? Check each point on the way to it
		slope := asteroid.SlopeTo(asteroid2)
		for point := asteroid.Add(slope); point != asteroid2; point = point.Add(slope) {
			if f.IsAsteroid(point) {
				// If a point on the way is an asteroid, nothing behind is visible
				return
			}
		}

		// If we got here we can see it
		asteroids = append(asteroids, asteroid2)
	})
	return asteroids
}

// Zaps all the asteroids and returns the order in which they were zapped
func (f *AsteroidField) ZapAsteroids(laserLocation Point2D) []Point2D {
	allZapped := make([]Point2D, 0)
	homeAngle := Point2D{0, -1}

	for {
		visible := f.VisibleAsteroids(laserLocation)
		if len(visible) == 0 {
			return allZapped
		}

		// Order them by angle
		sort.Slice(visible, func(i, j int) bool {
			vec1 := visible[i].Sub(laserLocation)
			vec2 := visible[j].Sub(laserLocation)

			return homeAngle.Angle(vec1) < homeAngle.Angle(vec2) 
		})

		for _, zapped := range(visible) {
			f.Field[zapped.Y][zapped.X] = false // zap
			allZapped = append(allZapped, zapped)
		}
	}
}

var Input = `#....#.....#...#.#.....#.#..#....#
#..#..##...#......#.....#..###.#.#
#......#.#.#.....##....#.#.....#..
..#.#...#.......#.##..#...........
.##..#...##......##.#.#...........
.....#.#..##...#..##.....#...#.##.
....#.##.##.#....###.#........####
..#....#..####........##.........#
..#...#......#.#..#..#.#.##......#
.............#.#....##.......#...#
.#.#..##.#.#.#.#.......#.....#....
.....##.###..#.....#.#..###.....##
.....#...#.#.#......#.#....##.....
##.#.....#...#....#...#..#....#.#.
..#.............###.#.##....#.#...
..##.#.........#.##.####.........#
##.#...###....#..#...###..##..#..#
.........#.#.....#........#.......
#.......#..#.#.#..##.....#.#.....#
..#....#....#.#.##......#..#.###..
......##.##.##...#...##.#...###...
.#.....#...#........#....#.###....
.#.#.#..#............#..........#.
..##.....#....#....##..#.#.......#
..##.....#.#......................
.#..#...#....#.#.....#.........#..
........#.............#.#.........
#...#.#......#.##....#...#.#.#...#
.#.....#.#.....#.....#.#.##......#
..##....#.....#.....#....#.##..#..
#..###.#.#....#......#...#........
..#......#..#....##...#.#.#...#..#
.#.##.#.#.....#..#..#........##...
....#...##.##.##......#..#..##....`