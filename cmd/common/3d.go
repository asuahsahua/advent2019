package common

import (
	"math"
)

type Point3D struct {
	X int
	Y int
	Z int
}

func Point3DStrs(xstr, ystr, zstr string) Point3D {
	return Point3D{
		X: LazyInt(xstr),
		Y: LazyInt(ystr),
		Z: LazyInt(zstr),
	}
}

func (p1 Point3D) Manhattan(p2 Point3D) int {
	return AbsI(p1.X-p2.X) + AbsI(p1.Y-p2.Y) + AbsI(p1.Z-p2.Z)
}

// Add a point to another, like vectors
func (p1 Point3D) Add(p2 Point3D) Point3D {
	return Point3D{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
		Z: p1.Z + p2.Z,
	}
}

func (p1 Point3D) Sub(p2 Point3D) Point3D {
	return Point3D{
		X: p1.X - p2.X,
		Y: p1.Y - p2.Y,
		Z: p1.Z - p2.Z,
	}
}

func (p1 Point3D) Dot(p2 Point3D) float64 {
	return float64(p1.X*p2.X + p1.Y*p2.Y + p1.Z*p2.Z)
}

func (p Point3D) Magnitude() float64 {
	return math.Sqrt(float64(p.X*p.X)) *
		math.Sqrt(float64(p.Y*p.Y)) *
		math.Sqrt(float64(p.Z*p.Z))
}

func (p Point3D) AbsSum() int {
	return AbsI(p.X) + AbsI(p.Y) + AbsI(p.Z)
}

func (p Point3D) fX() int {
	return p.X
}

func (p Point3D) fY() int {
	return p.Y
}

func (p Point3D) fZ() int {
	return p.Z
}
