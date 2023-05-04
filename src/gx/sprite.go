package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	I *Image
	T Transform
	Floating bool
}

func (s *Sprite) Draw(
	e *Engine,
	i *Image,
) {
	op := &ebiten.DrawImageOptions{}
	m := &Matrix{}

	m.Concat(s.T.Matrix(e))
	if e.camera != nil {
		m.Concat(e.camera.RealMatrix(
			e,
			true,
		))
	}

	op.GeoM = *m
	i.DrawImage(s.I, op)
}

