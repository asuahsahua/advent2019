package universe

import (
	. "github.com/asuahsahua/advent2019/cmd/common"
)

func (uni *Universe) CountTransfersBetween(body1Name string, body2Name string) int {
	body1 := uni.Objects[body1Name]
	PanicIf(body1 == nil, "Could not find celestial body by name: %s", body1Name)
	body2 := uni.Objects[body2Name]
	PanicIf(body2 == nil, "Could not find celestial body by name: %s", body2Name)

	// The number of transfers is equal to the number of non-common orbits!
	orbits1 := body1.Orbits()
	orbits2 := body2.Orbits()

	// Count up the common ones...
	common := 0
	maxCommons := MinI(len(orbits1), len(orbits2))
	for i := 0; i < maxCommons; i++ {
		if orbits1[i].Name == orbits2[i].Name {
			common++
		}
	}

	// And return the respective differences
	return len(orbits1) + len(orbits2) - 2 * common
}