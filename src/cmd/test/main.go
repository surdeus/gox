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
	ScaleSpeed gx.Float
}

type Debug struct{}

var (
	playerImg *gx.Image
)

func NewPlayer() *Player {
	return &Player{
		Sprite: &gx.Sprite{
			T: gx.Transform {
				S: gx.Vector{5, 5},
				RA: gx.Vector{320, 240},
			},
			I: playerImg,
			Visible: true,
		},
		MoveSpeed: 90.,
		ScaleSpeed: .2,
	}
}

func (p *Player) Start(e *gx.Engine) {
	c := e.Camera()
	c.T.RA = gx.V(360, -240)
}

func (p *Player) Update(e *gx.Engine) error {
	dt := e.DT()
	c := e.Camera()
	keys := e.Keys()

	for _, v := range keys {switch v {
	case ebiten.KeyArrowUp :
		c.T.P.Y += p.MoveSpeed * dt
	case ebiten.KeyArrowLeft :
		c.T.P.X -= p.MoveSpeed * dt
	case ebiten.KeyArrowDown :
		c.T.P.Y -= p.MoveSpeed * dt
	case ebiten.KeyArrowRight :
		c.T.P.X += p.MoveSpeed * dt
	case ebiten.KeyW :
		p.T.P.Y += p.MoveSpeed * dt
	case ebiten.KeyA :
		p.T.P.X -= p.MoveSpeed * dt
	case ebiten.KeyS :
		p.T.P.Y -= p.MoveSpeed * dt
	case ebiten.KeyD :
		p.T.P.X += p.MoveSpeed * dt
	case ebiten.KeyR :
		c.T.R += gx.Pi * p.ScaleSpeed * dt
	case ebiten.KeyT :
		c.T.R -= gx.Pi * p.ScaleSpeed * dt
	case ebiten.KeyF :
		if e.KeyIsPressed(ebiten.KeyShift) {
			c.T.S.X -= gx.Pi * p.ScaleSpeed * dt
		} else {
			c.T.S.X += gx.Pi * p.ScaleSpeed * dt
		}
	case ebiten.KeyG :
		if e.KeyIsPressed(ebiten.KeyShift) {
			c.T.S.Y -= gx.Pi * p.ScaleSpeed * dt
		} else {
			c.T.S.Y += gx.Pi * p.ScaleSpeed * dt
		}
	case ebiten.KeyZ :
		if e.KeyIsPressed(ebiten.KeyShift) {
			c.T.RA.X -= gx.Pi * p.MoveSpeed * dt
		} else {
			c.T.RA.X += gx.Pi * p.MoveSpeed * dt
		}
		log.Println(c.T.RA.X)
	case ebiten.KeyX :
		if e.KeyIsPressed(ebiten.KeyShift) {
			c.T.RA.Y -= gx.Pi * p.MoveSpeed * dt
		} else {
			c.T.RA.Y += gx.Pi * p.MoveSpeed * dt
		}
		log.Println(c.T.RA.Y)
	case ebiten.Key0 :
		e.Del(p)
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

func (d *Debug) IsVisible() bool {return true}

func main() {
	e := gx.New(&gx.WindowConfig{
		Title: "Test title",
		Width: 720,
		Height: 480,
		VSync: true,
	})

	var err error
	playerImg, err = gx.LoadImage(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}


	e.Add(0, NewPlayer())
	e.Add(1, &Debug{})
	e.Add(-1, gx.DrawableRectangle{
		Rectangle: gx.Rectangle{
			W: 100,
			H: 100,
			T: gx.T(),
		},
		Color: gx.Color{
			gx.MaxColorV,
			0,
			0,
			gx.MaxColorV,
		},
		Visible: true,
	})
	e.Run()
}
