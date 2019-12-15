package universe

import (
	"regexp"
	. "github.com/asuahsahua/advent2019/cmd/common"
)

type Universe struct {
	CenterOfMass *CelestialObject
	Objects      map[string]*CelestialObject
}

// Create a universe from the given orbits
func ParseUniverse(orbits string) *Universe {
	uni := &Universe{
		CenterOfMass: nil,
		Objects:      make(map[string]*CelestialObject, 0),
	}

	for _, orbit := range(SplitLines(orbits)) {
		uni.AddOrbitStr(orbit)
	}

	return uni
}

var orbitSpec *regexp.Regexp = regexp.MustCompile(`^(\w+)\)(\w+)$`)
func (uni *Universe) AddOrbitStr(spec string) {
	match := orbitSpec.FindStringSubmatch(spec)
	PanicIf(match == nil, "Could not parse orbiting spec line %q", spec)

	uni.AddOrbit(match[1], match[2])
}

func (uni *Universe) AddOrbit(centerName string, satelliteName string) {
	center := uni.GetCelestial(centerName)
	satellite := uni.GetCelestial(satelliteName)

	center.Satellites = append(center.Satellites, satellite)
	satellite.Orbiting = center
}

func (uni *Universe) GetCelestial(name string) *CelestialObject {
	if uni.Objects[name] == nil {
		uni.Objects[name] = NewCelestialObject(name)
	}

	return uni.Objects[name]
}

func (uni *Universe) SumOrbits() int {
	sum := 0
	for _, obj := range(uni.Objects) {
		sum += len(obj.Orbits())
	}
	return sum
}