package main

import (
	"github.com/surdeus/gox/src/gx"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"bytes"
	"log"
)

type Player struct {
	*gx.Sprite
}

var (
	playerImg *gx.Image
)

func NewPlayer() *Player {
	return &Player{
		Sprite: &gx.Sprite{
			Object: &gx.Object{
				T: gx.Transform {
					P: gx.Vector{1, 1},
					S: gx.Vector{1, 1},
				},
			},
			Image: playerImg,
		},
	}
}

func (p *Player) Update(e *gx.Engine) {
	p.Sprite.Object.T.P.Y += 0.1
}

func main() {
	e := gx.New(&gx.WindowConfig{
		Title: "Test title",
		Width: 480,
		Height: 320,
	})

	var err error
	playerImg, err = gx.LoadImage(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}


	e.Add(0, NewPlayer())
	e.Run()
}
