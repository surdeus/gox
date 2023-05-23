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
	// In fact are needed only to specify
	// relation of width and height.
	// Change transform to actually change things.
	W, H Float
	
}

// The type describes rectangle that can be drawn.
type DrawableRectangle struct {
	Rectangle
	Shader *Shader
	Color Color
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
	// Draw solid color if no shader.
	if r.Shader == nil {
		img := NewImage(1, 1)
		img.Set(0, 0, r.Color)
		
		t := r.T
		t.S.X *= r.W
		t.S.Y *= r.H
		
		rm := e.Camera().RealMatrix(e, true)
		m := t.Matrix(e)
		m.Concat(rm)
		opts := &ebiten.DrawImageOptions{
			GeoM: m,
		}
		i.DrawImage(img, opts)
		return
	}
	
	opts := &ebiten.DrawRectShaderOptions{
		GeoM: r.T.Matrix(e),
		Images: [4]*Image{
			NewImage(1, 1),
			nil,
			nil,
			nil,
		},
	}
	
	//w := int(r.W * r.T.S.X)
	//h := int(r.H * r.T.S.Y)
	
	i.DrawRectShader(1000, 1000, r.Shader, opts)
}

