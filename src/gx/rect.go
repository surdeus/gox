package gx

// The type describes rectangle geometry.
type Rect struct {
	// Position and rotate around.
	P, RA Point
	// Width and height.
	W, H Float
	// Rotation in radians.
	R Float
}

func (r Rect) ColliderSimplify() Rect {
	return r
}

