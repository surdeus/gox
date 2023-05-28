package gx

import (
	"math"
)

// The type represents mathematical equation of line and line itself.
type Line struct {
	K, C Float
}

type Liner interface {
	Line() Line
}

// The type represents a line segment.
type LineSegment [2]Point

// Check whether the liner is parallel to the line.
func (l Line) Parallel(liner Liner) bool {
	buf := liner.Line()
	
	if buf.K == l.K {
		return true
	}
	
	return false
}

// Get square of length of line segment.
func (ls LineSegment) LenSqr() Float {
	return Sqr(ls[0].X - ls[1].X) +
		Sqr(ls[0].Y - ls[1].Y)
}

// Get length of the line segment.
func (ls LineSegment) Len() Float {
	return math.Sqrt(ls.LenSqr())
}

func (l Line) Line() Line {
	return l
}

// Returns corresponding to the segment line line.
func (l LineSegment) Line() Line {
	p0 := l[0]
	p1 := l[1]
	
	k := (p0.Y - p1.Y) / (p0.X - p1.X)
	c := p0.Y - p0.X*k 
	
	return Line{k, c}
}


func (l LineSegment) Crosses(with any) (bool, Point) {
	switch with.(type) {
	case Line :
		return l.crossesLine(with.(Line))
	case LineSegment :
		return l.crossesLineSegment(with.(LineSegment))
	default:
		panic("The type that is not defined to be crossed")
	}
	
	
}

func (l LineSegment) crossesLineSegment(with LineSegment) (bool, Point) {
	return false, Point{}
}

func (l LineSegment) crossesLine(with Line) (bool, Point) {
	return false, Point{}
}

