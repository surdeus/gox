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
	
	Shader *Shader
	// Solid color of the rectangle.
	// Will be ignored if the Shader
	// field is not nil.
	Color Color
	
	Options ShaderOptions
	
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

func (r DrawableRectangle) IsVisible() bool {
	return r.Visible
}

func (r DrawableRectangle) Draw(
	e *Engine,
	i *Image,
) {
	t := r.T
	t.S.X *= r.W
	t.S.Y *= r.H
	
	rm := e.Camera().RealMatrix(e, true)
	m := t.Matrix(e)
	m.Concat(rm)
	
	// Draw solid color if no shader.
	if r.Shader == nil {
		img := NewImage(1, 1)
		img.Set(0, 0, r.Color)
		opts := &ebiten.DrawImageOptions{
			GeoM: m,
		}
		i.DrawImage(img, opts)
		return
	}
	
	// Use the Color as base image if no is provided.
	if r.Options.Images[0] == nil {
		r.Options.Images[0] = NewImage(1, 1)
		r.Options.Images[0].Set(0, 0, r.Color)
	} 
	
	// Drawing with shader.
	opts := &ebiten.DrawRectShaderOptions{
		GeoM: m,
		Images: r.Options.Images,
		Uniforms: r.Options.Uniforms,
	}
	
	w, h := r.Options.Images[0].Size()
	i.DrawRectShader(w, h, r.Shader, opts)
}

