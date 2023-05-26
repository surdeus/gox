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
	T Transform
	// Width and height.
	W, H Float
	
}

// The type describes rectangle that can be drawn.
type DrawableRectangle struct {
	Rectangle
	
	// Solid color of the rectangle.
	// Will be ignored if the Shader
	// field is not nil.
	Color Color
	
	// Shader to be applied
	Shader *Shader
	// Shader variables.
	Uniforms map[string] any
	// Shader images
	Images [4]*Image
	
	// Should be draw or not.
	Visible bool
}

// Return points of corners of the rectangle.
func (r Rectangle) Corners() []Point {
	return []Point{}
}

// Get 2 triangles that the rectangle consists of.
func (r Rectangle) Triangles() Triangles {
	return Triangles{}
}

/*func MustNewImage(w, h int) (*Image) {
	img, err := NewImage(w, h)
	if err != nil {
		panic(err)
	}
	
	return img
}*/

func NewImage(w, h int) (*Image) {
	return ebiten.NewImage(w, h)
}

func (r *DrawableRectangle) IsVisible() bool {
	return r.Visible
}

func (r *DrawableRectangle) Draw(
	e *Engine,
	i *Image,
) {
	t := r.T
	
	// Draw solid color if no shader.
	if r.Shader == nil {
		img := NewImage(1, 1)
		img.Set(0, 0, r.Color)
		
		t.S.X *= r.W
		t.S.Y *= r.H
		
		m := t.Matrix(e)
		rm := e.Camera().RealMatrix(e, true)
		
		m.Concat(rm)
		
		opts := &ebiten.DrawImageOptions{
			GeoM: m,
		}
		i.DrawImage(img, opts)
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

