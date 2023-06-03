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

// The type represents multiple line segments.
type LineSegments []LineSegment

// Check whether the liner is parallel to the line.
func (l Line) Parallel(liner Liner) bool {
	buf := liner.Line()
	
	if buf.K == l.K {
		return true
	}
	
	return false
}

func (l Line) Line() Line {
	return l
}

func (l Line) ContainsPoint(p Point) bool {
	buf := Line{0, p.Y}
	pc, ok := l.crossesLine(buf)
	if !ok {
		return false
	}
	
	return pc == p 
}

func (l Line) Crosses(with Liner) (Point, bool) {
	switch with.(type) {
	case Line :
		return l.crossesLine(with.(Line))
	case LineSegment :
		return with.(LineSegment).crossesLine(l)
	default:
		panic("unhandled type")
	}
}

func (l1 Line) crossesLine(l2 Line) (Point, bool) {
	if l1.Parallel(l2) {
		return Point{}, false
	}
	
	x := (l1.C - l2.C) / (l2.K - l1.K)
	y := l1.K*x + l1.C
	return Point{x, y}, true
}

func (l LineSegment) ContainsPoint(p Point) bool {
	line := l.Line()
	if !line.ContainsPoint(p) {
		return false
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

// Returns corresponding to the segment line line.
func (l LineSegment) Line() Line {
	p0 := l[0]
	p1 := l[1]
	
	k := (p0.Y - p1.Y) / (p0.X - p1.X)
	c := p0.Y - p0.X*k 
	
	return Line{k, c}
}

func (l LineSegment) Crosses(with Liner) (Point, bool) {
	switch with.(type) {
	case Line :
		return l.crossesLine(with.(Line))
	case LineSegment :
		return l.crossesLineSegment(with.(LineSegment))
	default:
		panic("The type that is not defined to be crossed")
	}
}

func (l LineSegment) Contains(what any) bool {
	switch what.(type) {
	case Point :
		return l.containsPoint(what.(Point))
	default :
		panic("Unexpected type")
	}
}

func (l LineSegment) containsPoint(p Point) bool {
	return false
}

func (l LineSegment) crossesLineSegment(with LineSegment) (Point, bool) {
	return Point{}, false
}

func (l LineSegment) crossesLine(with Line) (Point, bool) {
	return Point{}, false 
}

