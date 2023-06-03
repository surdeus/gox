package gx

// The type represents polygons.
// Fuck. The package gets too big.
// Should split it somehow.
type Polygon []Point

func (p Polygon) Edges() Edges {
	ret := Edges{}
	/*for i := range p {
		ret = append(ret, Edge{})
	}*/
	
	return ret
}

