package gx

type Float float64

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

