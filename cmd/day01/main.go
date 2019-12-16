package main

import (
	. "github.com/asuahsahua/advent2019/cmd/common"
)

// --- Day 1: The Tyranny of the Rocket Equation ---
// Santa has become stranded at the edge of the Solar System while delivering
// presents to other planets! To accurately calculate his position in space,
// safely align his warp drive, and return to Earth in time to save Christmas,
// he needs you to bring him measurements from fifty stars.
//
// Collect stars by solving puzzles. Two puzzles will be made available on each
// day in the Advent calendar; the second puzzle is unlocked when you complete
// the first. Each puzzle grants one star. Good luck!
//
// The Elves quickly load you into a spacecraft and prepare to launch.
//
// At the first Go / No Go poll, every Elf is Go until the Fuel Counter-Upper.
// They haven't determined the amount of fuel required yet.
func main() {
	// Fuel required to launch a given module is based on its mass.
	// Specifically, to find the fuel required for a module, take its mass,
	// divide by three, round down, and subtract 2.
	//
	// The Fuel Counter-Upper needs to know the total fuel requirement. To find
	// it, individually calculate the fuel needed for the mass of each module
	// (your puzzle input), then add together all the fuel values.
	//
	// What is the sum of the fuel requirements for all of the modules on your
	// spacecraft?
	masses := SplitInts(Input1, "\n")
	Part1("%d", SumModuleFuel(masses))

	// During the second Go / No Go poll, the Elf in charge of the Rocket
	// Equation Double-Checker stops the launch sequence. Apparently, you forgot
	// to include additional fuel for the fuel you just added.
	
	// Fuel itself requires fuel just like a module - take its mass, divide by
	// three, round down, and subtract 2. However, that fuel also requires fuel,
	// and that fuel requires fuel, and so on. Any mass that would require
	// negative fuel should instead be treated as if it requires zero fuel; the
	// remaining mass, if any, is instead handled by wishing really hard, which
	// has no mass and is outside the scope of this calculation.
	
	// So, for each module mass, calculate its fuel and add it to the total.
	// Then, treat the fuel amount you just calculated as the input mass and
	// repeat the process, continuing until a fuel requirement is zero or
	// negative.

	// What is the sum of the fuel requirements for all of the modules on your
	// spacecraft when also taking into account the mass of the added fuel?
	// (Calculate the fuel requirements for each module separately, then add
	// them all up at the end.)
	Part2("%d", SumModuleFuelFuel(masses))
}

func ModuleFuel(mass int) int {
	return (mass / 3) - 2
}

func SumModuleFuel(masses []int) int {
	sum := 0
	for _, mass := range(masses) {
		sum += ModuleFuel(mass)
	}
	return sum
}

func ModuleFuelFuel(mass int) int {
	fuel := ModuleFuel(mass)
	
	for nextFuel := ModuleFuel(fuel); nextFuel > 0; nextFuel = ModuleFuel(nextFuel) {
		fuel += nextFuel
	}
	
	return fuel
}

func SumModuleFuelFuel(masses []int) int {
	sum := 0
	for _, mass := range(masses) {
		sum += ModuleFuelFuel(mass)
	}
	return sum
}

var Input1 = `76542
97993
79222
55538
126710
77603
67546
129345
60846
52191
126281
85662
79245
78515
91236
126982
94593
63104
96955
122919
92047
63529
75949
65479
116132
55851
100051
120419
79243
109752
57719
131000
99825
92855
111945
58349
104867
53638
110072
111190
126422
72304
62865
113793
98395
86596
89219
135417
113665
87273
144161
97285
136308
79486
140622
138221
115714
142175
114524
50519
112963
109686
113104
50622
102019
96717
148433
70861
133918
89471
112281
109168
68965
109233
101051
52587
65339
97698
126416
61012
120883
81018
60398
112250
64253
98120
74640
134790
80984
61221
119815
96125
96105
87124
60042
141705
57290
57881
131585
51360`