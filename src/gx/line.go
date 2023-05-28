package gx

import (
	"math"
)

// The type represents mathematical equation of line.
type LineEquation struct {
	K, C Float
}

// The type represents a line segment.
type LineSegment [2]Point

// The type represents a line.
type Line LineSegment

// Get square of length of line segment.
func (ls LineSegment) LenSqr() Float {
	return Sqr(ls[0].X - ls[1].X) +
		Sqr(ls[0].Y - ls[1].Y)
}

// Get length of the line segment.
func (ls LineSegment) Len() Float {
	return math.Sqrt(ls.LenSqr())
}

// Returns corresponding to the line segment line.
func (ls LineSegment) Line() Line {
	return Line(ls)
}

func (l Line) Equation() LineEquation {
	p0 := l[0]
	p1 := l[1]
	
	k := (p0.Y - p1.Y) / (p0.X - p1.X)
	c := p0.Y - p0.X*k 
	
	return LineEquation{k, c}
}

