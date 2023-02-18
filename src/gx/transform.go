package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Matrix = ebiten.GeoM
type Float = float64

type Vector struct {
	X, Y Float
}

type Transform struct {
	// Position, scale, rotate around.
	P, S, RA Vector
	// Rotation angle in radians.
	R Float
}

func T() Transform {
	ret := Transform{}
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

