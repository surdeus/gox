package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
	textx "github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"io"
	"os"
	//"fmt"
)

type Font = font.Face

// The type represents drawable text.
type Text struct {
	Transform
	ShaderOptions
	Colority
	Visibility
	
	Font Font
	
	// Text to be drawn.
	Text string
}

func LoadFont(fp string, size Float) (Font, error) {
	nilRet := Font(nil)
	
	f, err := os.Open(fp)
	if err != nil {
		return nilRet, err
	}
	
	dat, err := io.ReadAll(f)
	if err != nil {
		return nilRet, err
	}
	
	tt, err := opentype.Parse(dat)
	if err != nil {
		return nilRet, err
	}
	
	ret, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size: size,
		DPI: 96,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nilRet, err
	}
	
	return ret, nil
}

func (text *Text) Draw(e *Engine, i *Image) {
	var opts ebiten.DrawImageOptions
	opts.ColorM.ScaleWithColor(text.Color)
	opts.GeoM = text.Matrix()
	textx.DrawWithOptions(i, text.Text, text.Font, &opts)
}

