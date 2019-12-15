package universe

// Attempting to treat these celestials as a Directed Acyclical Graph (DAG).
// Things are going to get really nasty if that assumption is ever broken and
// there are cycles...
type CelestialObject struct {
	Name string

	Orbiting *CelestialObject
	Satellites []*CelestialObject

	DirectSatellitesCount int
	IndirectSatellitesCount int
}

func NewCelestialObject(name string) *CelestialObject {
	return &CelestialObject{
		Name: name,
		Orbiting: nil,
		Satellites: make([]*CelestialObject, 0),
	}
}

func (obj *CelestialObject) Orbits() []*CelestialObject {
	// If we're the root, then we have no orbits
	if obj.Orbiting == nil {
		return []*CelestialObject{}
	}

	// Otherwise, we have the parent and the orbits of the parent
	parentOrbits := obj.Orbiting.Orbits()
	// Heh, this might get me into trouble! This is going to append
	// the Orbiting celestial to the inner array of the slice returned
	// from the parent's Orbits(). 
	// This may be okay if we [[ never, ever, ever, EVER, EVERR modify the array contents after this ]]
	return append(parentOrbits, obj.Orbiting)
}