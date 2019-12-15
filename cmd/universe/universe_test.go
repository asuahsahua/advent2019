package universe

import (
	"testing"
	. "github.com/stretchr/testify/assert"
)

// This file holds nonspecific tests from each day's prompts

// --- Day 5 - Part 2 ---
func TestOrbits(t *testing.T) {
	// For example, suppose you have the following map:
	universe := ParseUniverse(example05_1)

	// Visually, the above map of orbits looks like this:
	//         G - H       J - K - L
	//        /           /
	// COM - B - C - D - E - F
	//                \
	//                 I

	// In this visual representation, when two objects are connected by a line,
	// the one on the right directly orbits the one on the left. Here, we can
	// count the total number of orbits as follows:

	//  * D directly orbits C and indirectly orbits B and COM, a total of 3 orbits.
	Equal(t, 3, len(universe.GetCelestial("D").Orbits()))
	//  * L directly orbits K and indirectly orbits J, E, D, C, B, and COM, a total of 7 orbits.
	Equal(t, 7, len(universe.GetCelestial("L").Orbits()))
	//  * COM orbits nothing.
	Equal(t, 0, len(universe.GetCelestial("COM").Orbits()))

	// The total number of direct and indirect orbits in this example is 42.
	Equal(t, 42, universe.SumOrbits())
}

func TestCountTransfersBetween(t *testing.T) {
	// For example, suppose you have the following map:
	universe := ParseUniverse(example05_2)

	// Visually, the above map of orbits looks like this:

	//                           YOU
	//                          /
	//         G - H       J - K - L
	//        /           /
	// COM - B - C - D - E - F
	//                \
	//                 I - SAN

	// In this example, YOU are in orbit around K, and SAN is in orbit around I.
	// To move from K to I, a minimum of 4 orbital transfers are required:

	//     K to J
	//     J to E
	//     E to D
	//     D to I

	Equal(t, 4, universe.CountTransfersBetween("YOU", "SAN"))
}

var example05_1 string = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`

var example05_2 string = `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`