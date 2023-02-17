package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	*Object
	*Image
}

func (s *Sprite) Draw(
	e *Engine,
	i *Image,
) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM = s.Object.T.Matrix()

	i.DrawImage(s.Image, op)
}

