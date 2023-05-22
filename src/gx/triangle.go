package gx

import (
	"math"
)

// The structure of a triangle. What more you want to hear?
type Triangle [3]Point
type Triangles []Triangle

// Returns the area of the triangle.
func (t Triangle) Area() Float {
	x1 := t[0].X
	y1 := t[0].Y
	
	x2 := t[1].X
	y2 := t[1].Y
	
	x3 := t[2].X
	y3 := t[2].Y
	
	return math.Abs( (x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2))/2)
}

