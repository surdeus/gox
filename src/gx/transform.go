package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
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

func V(x, y Float) Vector {
	return Vector{x, y}
}

// Returns the vector with the matrix applied
func (v Vector) Apply(m *Matrix) Vector {
	x, y := m.Apply(v.X, v.Y)
	return V(x, y)
}

// Returns the vector rotated by "a" angle in radians.
func (v Vector) Rotate(a Float) Vector {
	m := &Matrix{}
	m.Rotate(a)
	return v.Apply(m)
}

// Returns the normalized vector.
func (v Vector) Norm() Vector {
	l := math.Sqrt(v.X*v.X + v.Y*v.Y)
	return V(v.X / l, v.Y / l)
}

func T() Transform {
	ret := Transform{}
	return ret
}

// Returns the GeoM with corresponding
// to the transfrom transformation 
func (t Transform)Matrix(e *Engine) Matrix {
	g := &Matrix{}

	g.Scale(t.S.X, t.S.Y)
	g.Translate(-t.RA.X, -t.RA.Y)
	g.Rotate(t.R)
	g.Translate(t.P.X, -t.P.Y)

	return *g
}

