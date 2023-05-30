package gx

import (
	//"github.com/hajimehoshi/ebiten/v2"
	//"math"
)

type Transform struct {
	// Position, scale, rotate around(relatively of position, not absolute).
	P, S, RA Vector
	// Rotation angle in radians.
	R Float
}

// Returns empty Transform.
func T() Transform {
	ret := Transform{
		S: Vector{1, 1},
	}
	return ret
}

// Returns the GeoM with corresponding
// to the transfrom transformation 
func (t Transform)Matrix() Matrix {
	g := &Matrix{}

	g.Scale(t.S.X, t.S.Y)
	g.Translate(-t.RA.X, -t.RA.Y)
	g.Rotate(t.R)
	g.Translate(t.P.X, -t.P.Y)

	return *g
}

