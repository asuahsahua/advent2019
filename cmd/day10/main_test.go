package main

import (
	"testing"
	. "github.com/stretchr/testify/assert"
	"github.com/asuahsahua/advent2019/cmd/common"
)

func TestExamples(t *testing.T) {
	Equal(t, 8, BestAsteroidScoreStr(example1))
	Equal(t, 33, BestAsteroidScoreStr(example2))
	Equal(t, 35, BestAsteroidScoreStr(example3))
	Equal(t, 41, BestAsteroidScoreStr(example4))
	Equal(t, 210, BestAsteroidScoreStr(example5))
}

func TestZapAsteroids(t *testing.T) {
	t.Skip("Solution NYI")

	field := NewAsteroidFieldStr(example5)
	zapped := field.ZapAsteroids(common.Point2D{11, 13})

    // The 1st asteroid to be vaporized is at 11,12.
	Equal(t, common.Point2D{11, 12}, zapped[0])
    // The 2nd asteroid to be vaporized is at 12,1.
	Equal(t, common.Point2D{12, 1}, zapped[1])
    // The 3rd asteroid to be vaporized is at 12,2.
	Equal(t, common.Point2D{12, 2}, zapped[2])
    // The 10th asteroid to be vaporized is at 12,8.
	Equal(t, common.Point2D{12, 8}, zapped[9])
    // The 20th asteroid to be vaporized is at 16,0.
	Equal(t, common.Point2D{16, 8}, zapped[19])
    // The 50th asteroid to be vaporized is at 16,9.
	Equal(t, common.Point2D{16, 9}, zapped[49])
    // The 100th asteroid to be vaporized is at 10,16.
	Equal(t, common.Point2D{10, 16}, zapped[99])
    // The 199th asteroid to be vaporized is at 9,6.
	Equal(t, common.Point2D{9, 6}, zapped[198])
    // The 200th asteroid to be vaporized is at 8,2.
	Equal(t, common.Point2D{8, 2}, zapped[199])
    // The 201st asteroid to be vaporized is at 10,9.
	Equal(t, common.Point2D{8, 2}, zapped[201])
    // The 299th and final asteroid to be vaporized is at 11,1.
	Equal(t, common.Point2D{8, 2}, zapped[298])
	Equal(t, 299, len(zapped))
}

func TestConfirmedAnswer(t *testing.T) {
	Equal(t, 267, BestAsteroidScoreStr(Input))
}

var example1 = `.#..#
.....
#####
....#
...##`

var example2 = `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`

var example3 = `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`

var example4 = `.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`

var example5 = `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`

var example6 = `.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....#...###..
..#.#.....#....##`