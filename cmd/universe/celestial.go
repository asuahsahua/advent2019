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

// Recursive!
func (obj *CelestialObject) TotalOrbitCount() int {
	if obj.Orbiting == nil {
		return 0
	}

	return obj.Orbiting.TotalOrbitCount() + 1
}