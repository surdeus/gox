package gx

// The structure represents any circles.
type Circle struct {
	// Position.
	P Point
	// Radius.
	R Float
}

func (c Circle) ColliderSimplify() Rectangle {
	return Rectangle{
		W: c.R * 2,
		H: c.R * 2,
	}
}

