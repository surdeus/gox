package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"io"
	"math"
)

type Image = ebiten.Image

type ColorV uint32
type ColorM = ebiten.ColorM
type Color struct {
	R, G, B, A ColorV
}

type Colority struct {
	Color Color
}

type Visibility struct {
	Visible bool
}

// The interface describes anything that can be
// drawn. It will be drew corresponding to
// the layers order.
type Drawer interface {
	Draw(*Engine, *Image)
}

type Visibler interface {
	IsVisible() bool
}

const (
	MaxColorV = math.MaxUint32
)

func (v Visibility) IsVisible() bool {
	return v.Visible
}

func LoadImage(input io.Reader) (*Image, error) {
	img, _, err := image.Decode(input)
	if err != nil {
		return nil, err
	}

	ret := ebiten.NewImageFromImage(img)
	return ret, nil
}

func NewImage(w, h int) (*Image) {
	return ebiten.NewImage(w, h)
}


func (c Color) RGBA() (r, g, b, a uint32) {
	return uint32(c.R), uint32(c.G), uint32(c.B), uint32(c.A)
}

