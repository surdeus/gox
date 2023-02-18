package main

import (
	"github.com/surdeus/gox/src/gx"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2"
	"bytes"
	"log"
	"strings"
)

type Player struct {
	*gx.Sprite
	MoveSpeed gx.Float
}

type Debug struct{}

var (
	playerImg *gx.Image
)

func NewPlayer() *Player {
	return &Player{
		Sprite: &gx.Sprite{
			Object: &gx.Object{
				T: gx.Transform {
					P: gx.Vector{100, 150},
					S: gx.Vector{5, 5},
					RA: gx.Vector{200, 200},
				},
			},
			Image: playerImg,
		},
		MoveSpeed: 90.,
	}
}

func (p *Player) Update(e *gx.Engine) error {
	dt := e.DT()
	c := e.Camera()
	keys := e.Keys()

	for _, v := range keys {switch v {
	case ebiten.KeyArrowUp :
		c.Object.T.P.Y += p.MoveSpeed * dt
	case ebiten.KeyArrowLeft :
		c.Object.T.P.X -= p.MoveSpeed * dt
	case ebiten.KeyArrowDown :
		c.Object.T.P.Y -= p.MoveSpeed * dt
	case ebiten.KeyArrowRight :
		c.Object.T.P.X += p.MoveSpeed * dt
	case ebiten.KeyW :
		p.Object.T.P.Y += p.MoveSpeed * dt
	case ebiten.KeyA :
		p.Object.T.P.X -= p.MoveSpeed * dt
	case ebiten.KeyS :
		p.Object.T.P.Y -= p.MoveSpeed * dt
	case ebiten.KeyD :
		p.Object.T.P.X += p.MoveSpeed * dt
	case ebiten.KeyR :
		c.Object.T.R += gx.Pi * .3 * dt
	case ebiten.KeyT :
		c.Object.T.R -= gx.Pi * .3 * dt
	}}

	return nil
}

func (d *Debug) Draw(
	e *gx.Engine,
	i *gx.Image,
) {
	keyStrs := []string{}
	for _, k := range e.Keys() {
		keyStrs = append(keyStrs, k.String())
	}
	e.DebugPrint(i, strings.Join(keyStrs, ", "))
}

func main() {
	e := gx.New(&gx.WindowConfig{
		Title: "Test title",
		Width: 720,
		Height: 480,
	})

	var err error
	playerImg, err = gx.LoadImage(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}


	e.Add(0, NewPlayer())
	e.Add(1, &Debug{})
	e.Run()
}
