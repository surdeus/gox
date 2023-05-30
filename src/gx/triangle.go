package gx

import (
	"math"
)

// The structure of a triangle. What more you want to hear?
type Triangle [3]Point
type Triangles []Triangle
type DrawableTriangle struct {
	Triangle
	ShaderOptions
}

// Returns the area of the triangle.
func (t Triangle) Area() Float {
	x1 := t[0].X
	y1 := t[0].Y
	
	x2 := t[1].X
	y2 := t[1].Y
	
	x3 := t[2].X
	y3 := t[2].Y
	
	return math.Abs( (x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2))/2)
}
func sqr(v Float) Float {
	return v * v
}
// Return squares of lengths of sides of the triangle.
func (t Triangle) SideLengthSquares() ([3]Float) {
	
	l1 := LineSegment{t[0], t[1]}.LenSqr()
	l2 := LineSegment{t[1], t[2]}.LenSqr()
	l3 := LineSegment{t[2], t[0]}.LenSqr()
	
	return [3]Float{l1, l2, l3}
}

// Check whether the point is in the triangle.
func (t Triangle) ContainsPoint(p Point) bool {
	d1 := Triangle{p, t[0], t[1]}.Sgn()
	d2 := Triangle{p, t[1], t[2]}.Sgn()
	d3 := Triangle{p, t[2], t[0]}.Sgn()
	
	neg := (d1 < 0) || (d2 < 0) || (d3 < 0)
	pos := (d1 > 0) || (d2 > 0) || (d3 > 0)
	
	return !(neg && pos)
}

func (t Triangle) Sgn() Float {
	return (t[0].X - t[2].X) * (t[1].Y - t[2].Y)  -
		(t[1].X - t[2].X) * (t[0].Y - t[2].Y)
}

func (ts Triangles) ContainsPoint(p Point) bool {
	for _, t := range ts {
		if t.ContainsPoint(p) {
			return true
		}
	}
	
	return false
}

//func (t Triangle)
/*
func (r *DrawableRectangle) Draw(
	e *Engine,
	i *Image,
) {
	t := r.T
	
	// Draw solid color if no shader.
	if r.Shader == nil {
		t.S.X *= r.W
		t.S.Y *= r.H
		
		m := t.Matrix(e)
		rm := e.Camera().RealMatrix(e, true)
		
		m.Concat(rm)
		
		opts := &ebiten.DrawImageOptions{
			GeoM: m,
		}
		i.DrawTriangles(img, opts)
		return
	}
	
	// Use the Color as base image if no is provided.
	var did bool
	if r.Images[0] == nil {
		r.Images[0] = NewImage(1, 1)
		r.Images[0].Set(0, 0, r.Color)
		did = true
	} 
	
	w, h := r.Images[0].Size()
	if !did {
		t.S.X /= Float(w)
		t.S.Y /= Float(h)
		
		t.S.X *= r.W
		t.S.Y *= r.H
	}
	
	
	rm := e.Camera().RealMatrix(e, true)
	m := t.Matrix(e)
	m.Concat(rm)
	
	// Drawing with shader.
	opts := &ebiten.DrawRectShaderOptions{
		GeoM: m,
		Images: r.Images,
		Uniforms: r.Uniforms,
	}
	i.DrawRectShader(w, h, r.Shader, opts)
}
*/

