package gx

import (
	//"github.com/hajimehoshi/ebiten/v2"
	"math"
)

type Vector struct {
	X, Y Float
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
