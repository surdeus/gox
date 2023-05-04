package gx

import (
	"math"
)

// Implements the camera component
// for the main window.
type Camera struct {
	T Transform
}

const (
	Pi = math.Pi
)

// Returns the matrix satysfying camera
// position, scale and rotation to apply
// it to the objects to get the real
// transform to display on the screen.
// (Should implement buffering it so we do not
//  need to calculate it each time for each object. )
func (c *Camera)RealMatrix(
	e *Engine,
	scale bool,
) Matrix {
	g := &Matrix{}


	if scale {
		g.Scale(
			c.T.S.X,
			c.T.S.Y,
		)
	}

	g.Translate(
		-c.T.P.X,
		c.T.P.Y,
	)
	g.Rotate(-c.T.R)

	g.Translate(
		c.T.RA.X,
		-c.T.RA.Y,
	)

	return *g
}

