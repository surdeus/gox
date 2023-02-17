package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"io"
)

type Image = ebiten.Image
type Rectangle = image.Rectangle
type Point = image.Point

// The interface describes anything that can be
// drawn. It will be drew corresponding to
// the layers order.
type Drawer interface {
	Draw(*Engine, *Image)
}

func LoadImage(input io.Reader) (*Image, error) {
	img, _, err := image.Decode(input)
	if err != nil {
		return nil, err
	}

	ret := ebiten.NewImageFromImage(img)
	return ret, nil
}

