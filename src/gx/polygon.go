package gx

// The type represents polygons.
// Fuck. The package gets too big.
// Should split it somehow.
type Polygon []Point

// Returns slice of edges.
func (p Polygon) Edges() Edges {
	ret := Edges{}
	
	l := p.Len()
	for i := range p {
		ret = append(ret, Edge{p[i], p[(i+1)%l]})
	}
	
	return ret
}

func (p Polygon) Len() int {
	return len(p)
}

func (p Polygon) Vertices() []Vertex {
	return p
}

func (p Polygon) Barycenter() Point {
	ret := Point{}
	for _, v := range p {
		ret.X += v.X
		ret.Y += v.Y
	}
	
	l := Float(len(p))
	ret.X /= l
	ret.Y /= l
	
	return ret
}

/*
func (p Polygon) anyPointInside() Point {
	edges := p.Edges()
	for _, e := range edges {
		if
	}
}
*/

func (p Polygon) Triangles() Triangles {
	/*
	if len(p) < 3 {
		return Triangles{}
	} else 	
	vertices = p.Vertices()
	ret := Triangles{}
	
	i := 0
	for len(vertices) != 3{

		i1 = i % len(vertices)
		i2 = (i+1) % len(vertices)
		i3 = (i+2) % len(vertices)
		
		l1 = LineSegment{vertices[i1], vertices[i2]}.Line()
		l2 = LineSegment{vertices[i1], vertices[i3]}.Line()
		
		if l1.K > 0 {
			if l1.K 
		} else {
		}
	}
	
	return append(
		ret,
		Triangle{p[0], p[1], p[2]},
	)
	*/
	
	return Triangles{}
}

