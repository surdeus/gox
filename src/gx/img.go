package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"io"
	"math"
)

type ColorV uint32
type Image = ebiten.Image
type ColorM = ebiten.ColorM
type Color struct {
	R, G, B, A ColorV
}

// The interface describes anything that can be
// drawn. It will be drew corresponding to
// the layers order.
type Drawer interface {
	Draw(*Engine, *Image)
	IsVisible() bool
}

const (
	MaxColorV = math.MaxUint32
)

func LoadImage(input io.Reader) (*Image, error) {
	img, _, err := image.Decode(input)
	if err != nil {
		return nil, err
	}

	ret := ebiten.NewImageFromImage(img)
	return ret, nil
}

func (c Color) RGBA() (r, g, b, a uint32) {
	return uint32(c.R), uint32(c.G), uint32(c.B), uint32(c.A)
}

