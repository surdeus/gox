package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	*Object
	*Image
	Floating bool
}

func (s *Sprite) Draw(
	e *Engine,
	i *Image,
) {
	op := &ebiten.DrawImageOptions{}
	m := &Matrix{}

	m.Concat(s.Object.T.Matrix(e))
	if e.camera != nil {
		m.Concat(e.camera.RealMatrix(
			e,
			true,
		))
	}

	op.GeoM = *m
	i.DrawImage(s.Image, op)
}

