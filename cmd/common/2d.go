package common 

import (
	"fmt"
	"math"
)

type Point2D struct{
	X int
	Y int
}

func Point2DI(x, y int) Point2D {
	return Point2D {
		X: x,
		Y: y,
	}
}

func (p1 Point2D) Manhattan(p2 Point2D) int {
	return AbsI(p1.X - p2.X) + AbsI(p1.Y - p2.Y)
}

// Determines the slope from one to the other as a reduced fraction
func (p1 Point2D) SlopeTo(p2 Point2D) Point2D {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y 

	gcd := GCD(dx, dy)
	if gcd < 0 {
		// we _need_ a positive gcd
		gcd = -1 * gcd
	}

	return Point2D{
		X: dx / gcd,
		Y: dy / gcd,
	}
}

// Add a point to another, like vectors
func (p1 Point2D) Add(p2 Point2D) Point2D {
	return Point2D {
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
	}
}

func (p1 Point2D) Sub(p2 Point2D) Point2D {
	return Point2D {
		X: p1.X - p2.X,
		Y: p1.Y - p2.Y,
	}
}

func (p1 Point2D) Dot(p2 Point2D) float64 {
	return float64(p1.X * p2.X + p1.Y * p2.Y)
}

func (p Point2D) Magnitude() float64 {
	return math.Sqrt(float64(p.X * p.X)) * math.Sqrt(float64(p.Y * p.Y))
}

// Gets the angle from Vector1 to Vector2 in radians
func (p1 Point2D) Angle(p2 Point2D) float64 {
	// cos(th) = u.v/(|u||v|)
	cosT := p1.Dot(p2) / (p1.Magnitude() * p2.Magnitude())
	return math.Acos(cosT)
}

// CardinalAdjacents() returns the points cardinally adjacent to the point
func (p1 Point2D) CardinalAdjacents() []Point2D {
	return []Point2D{
		p1.Add(Point2D{X: 1, Y: 0}),
		p1.Add(Point2D{X: 0, Y: 1}),
		p1.Add(Point2D{X: -1, Y: 0}),
		p1.Add(Point2D{X: 0, Y: -1}),
	}
}

type BoundingBox struct{
	MinX int
	MaxX int
	MinY int
	MaxY int
}

func ResolveBoundingBox(points []Point2D) BoundingBox {
	box := BoundingBox{
		MinX: MAX_INT64,
		MaxX: MIN_INT64,
		MinY: MAX_INT64,
		MaxY: MIN_INT64,
	}

	for _, point := range points {
		if point.X > box.MaxX {
			box.MaxX = point.X
		}
		if point.X < box.MinX {
			box.MinX = point.X
		}
		if point.Y > box.MaxY {
			box.MaxY = point.Y
		}
		if point.Y < box.MinY {
			box.MinY = point.Y
		}
	}

	return box
}

func (b BoundingBox) PrintPoints(cb func(Point2D) byte) {
	for x := b.MinX; x <= b.MaxX; x++ {
		for y := b.MinY; y <= b.MaxY; y++ {
			fmt.Printf("%c", cb(Point2D{X: x, Y: y}))
		}
		fmt.Printf("\n")
	}
}