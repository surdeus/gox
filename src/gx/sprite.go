package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Transform
	ShaderOptions
	Floating, Visible bool
}

func (s *Sprite) Draw(
	e *Engine,
	i *Image,
) {
	// Nothing to draw.
	if s.Images[0] == nil {
		return
	}
	
	m := &Matrix{}

	m.Concat(s.Matrix())
	if !s.Floating {
		m.Concat(e.Camera().RealMatrix(
			e,
		))
	}

	// Drawing without shader.
	if s.Shader == nil {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM = *m
		i.DrawImage(s.Images[0], opts)
		return
	}
	
	// Drawing with shader.
	w, h := s.Images[0].Size()
	opts := &ebiten.DrawRectShaderOptions{
		Images: s.Images,
		Uniforms: s.Uniforms,
		GeoM: *m,
	}
	i.DrawRectShader(w, h, s.Shader, opts)
}

func (s *Sprite) IsVisible() bool {
	return s.Visible
}

