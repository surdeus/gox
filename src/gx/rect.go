package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/vector"
	//"fmt"
	//"image"
)

// The type describes rectangle geometry.
type Rectangle struct {
	// Position of up left corner
	// and the point to
	// rotate around(relatively of position, not absolute).
	// Scale represent width and height.
	Transform
}

// The type describes rectangle that can be drawn.
type DrawableRectangle struct {
	Rectangle
	ShaderOptions	
	
	// Solid color of the rectangle.
	// Will be ignored if the Shader
	// field is not nil.
	Color Color
	
	// Should be draw or not.
	Visible bool
}

// Return points of corners of the rectangle.
func (r Rectangle) Corners() []Point {
	return []Point{}
}

// Get 2 triangles that the rectangle consists of.
func (r Rectangle) Triangles() Triangles {
	m := r.Matrix()
	
	p1 := V(0, 0).Apply(&m)
	p2 := V(1, 0).Apply(&m)
	p3 := V(1, -1).Apply(&m)
	p4 := V(0, -1).Apply(&m)
	
	//fmt.Println("in:", p1, p2, p3, p4)
	
	return Triangles{
		Triangle{p1, p2, p3},
		Triangle{p1, p4, p3},
	}
}

// Check whether the rectangle contains the point.
func (r Rectangle) ContainsPoint(p Point) bool {
	return r.Triangles().ContainsPoint(p)
}

// Check whether the drawable rectangle should be drawn.
func (r *DrawableRectangle) IsVisible() bool {
	return r.Visible
}

func (r *DrawableRectangle) Draw(
	e *Engine,
	i *Image,
) {
	t := r.Transform
	
	// Draw solid color if no shader.
	if r.Shader == nil {
		img := NewImage(1, 1)
		img.Set(0, 0, r.Color)
		
		m := t.Matrix()
		rm := e.Camera().RealMatrix(e)
		
		m.Concat(rm)
		
		opts := &ebiten.DrawImageOptions{
			GeoM: m,
		}
		i.DrawImage(img, opts)
		return
	}
	
	
	// Use the Color as base image if no is provided.
	//var did bool
	if r.Images[0] == nil {
		r.Images[0] = NewImage(1, 1)
		r.Images[0].Set(0, 0, r.Color)
		//did = true
	} 
	
	w, h := r.Images[0].Size()
	/*if !did {
		t.S.X /= Float(w)
		t.S.Y /= Float(h)
		
		t.S.X *= r.W
		t.S.Y *= r.H
	}*/
	
	
	rm := e.Camera().RealMatrix(e)
	m := t.Matrix()
	m.Concat(rm)
	
	// Drawing with shader.
	opts := &ebiten.DrawRectShaderOptions{
		GeoM: m,
		Images: r.Images,
		Uniforms: r.Uniforms,
	}
	i.DrawRectShader(w, h, r.Shader, opts)
}

