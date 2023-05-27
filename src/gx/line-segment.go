package gx

import (
	"math"
)

type LineSegment [2]Point

func (ls LineSegment) LenSqr() Float {
	return Sqr(ls[0].X - ls[1].X) +
		Sqr(ls[0].Y - ls[1].Y)
}

func (ls LineSegment) Len() Float {
	return Sqrt(ls.LenSqr())
}

