package gx

import (
	"math"
)

// Implements the camera component
// for the main window.
type Camera struct {
	*Object
}

const (
	Pi = math.Pi
)

// Returns the matrix satysfying camera
// position, scale and rotation to apply
// it to the objects to get the real
// transform to display on the screen.
func (c *Camera)RealMatrix(
	e *Engine,
	scale bool,
) Matrix {
	g := &Matrix{}


	if scale {
		g.Scale(
			c.Object.T.S.X,
			c.Object.T.S.Y,
		)
	}

	g.Translate(
		-c.Object.T.P.X,
		c.Object.T.P.Y,
	)
	g.Rotate(-c.Object.T.R)

	g.Translate(
		c.Object.T.RA.X,
		-c.Object.T.RA.Y,
	)

	return *g
}

