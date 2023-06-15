package gx

type PolygonTriangle struct {
	P, S int
}

type Polygon struct {
	Transform
	base Triangle
	triangles []PolygonTriangle
}

type DrawablePolygon struct {
	Polygon
	ShaderOptions
}

func NewPolygon(base Triangle) *Polygon {
	ret := &Polygon{
		Transform: T(),
		base: base,
	}
	return ret
}

func (p *Polygon) Triangles() Triangles {
	return Triangles{}
}

