package gx

type Circle struct {
	// Position
	P Point
	// Radius
	R Float
}

func (c Circle) ColliderSimplify() Rect {
	return Rect{
		W: c.R * 2,
		H: c.R * 2,
	}
}

