package main

import (
	"github.com/surdeus/gox/src/gx"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2"
	"bytes"
	"log"
	"strings"
	"fmt"
	"math/rand"
)

type Player struct {
	*gx.Sprite
	MoveSpeed gx.Float
	ScaleSpeed gx.Float
}

type Debug struct{}

type Rect struct {
	*gx.DrawableRectangle
}

type Tri struct {
	*gx.DrawablePolygon
}

type MyText struct {
	*gx.Text
}

func NewTri() *Tri {
	ret := &Tri{}
	ret.DrawablePolygon = &gx.DrawablePolygon{}
	ret.Transform.S = gx.V(1, 1)
	
	ret.Triangles = gx.Triangles{
		gx.Triangle{
			gx.V(0, 0),
			gx.V(100, 100),
			gx.V(0, -50),
		},
		gx.Triangle{
			gx.V(0, 0),
			gx.V(-100, -100),
			gx.V(0, 50),
		},
	}
	ret.Color = gx.Color{gx.MaxColorV, gx.MaxColorV, 0, gx.MaxColorV}
	ret.Visible = true
	
	return ret
}

func NewRect() *Rect {
	ret := &Rect{&gx.DrawableRectangle{
			Rectangle: gx.Rectangle{
				Transform: gx.Transform{
					S: gx.Vector{
						X: 200,
						Y: 400,
					},
				},
			},
			/*Shader: gx.SolidWhiteColorShader,
			Options: gx.ShaderOptions{
				Images: [4]*gx.Image{
					playerImg,
					nil,
					nil,
					nil,
				},
			},*/
	}}
	
	ret.Visible = true
	ret.Color = gx.Color{
		gx.MaxColorV,
		0,
		0,
		gx.MaxColorV,
	}
	
	return ret
}

func (r *Rect) Update(e *gx.Engine) error {
	//r.R += 0.3 * e.DT()
	return nil
}

var (
	playerImg *gx.Image
	player *Player
	rectMove gx.Rectangle
	rect *Rect
	tri *Tri
)

func NewPlayer() *Player {
	ret := &Player{
		Sprite: &gx.Sprite{
			Transform: gx.Transform {
				S: gx.Vector{5, 5},
				RA: gx.Vector{.5, .5},
			},
			Visible: true,
			ShaderOptions: gx.ShaderOptions {
				Shader: gx.SolidWhiteColorShader,
				Uniforms: make(map[string] any),
			},
		},
		MoveSpeed: 90.,
		ScaleSpeed: .2,
	}
	
	ret.Images[0] = playerImg
	
	return ret
}

func (p *Player) Draw(e *gx.Engine, i *gx.Image) {
	p.Sprite.Draw(e, i)
	t := p.Transform
	t.S.X *= 4.
	t.S.Y *= 4.
	rectMove = gx.Rectangle{
		Transform: t,
	}
	r := &gx.DrawableRectangle{
		Rectangle: rectMove,
	}
	r.Color = gx.Color{0, 0, gx.MaxColorV, gx.MaxColorV}
	r.Draw(e, i)
}

func (p *Player) Start(e *gx.Engine, v ...any) {
	fmt.Println("starting")
	c := e.Camera()
	c.RA = gx.V(360, 240)
}

func (p *Player) Update(e *gx.Engine) error {
	dt := e.DT()
	c := e.Camera()
	keys := e.Keys()
	
	p.Uniforms["Random"] = any(rand.Float32())
	for _, v := range keys {switch v {
	case ebiten.KeyArrowUp :
		c.P.Y += p.MoveSpeed * dt
	case ebiten.KeyArrowLeft :
		c.P.X -= p.MoveSpeed * dt
	case ebiten.KeyArrowDown :
		c.P.Y -= p.MoveSpeed * dt
	case ebiten.KeyArrowRight :
		c.P.X += p.MoveSpeed * dt
	case ebiten.KeyW :
		p.P.Y += p.MoveSpeed * dt
	case ebiten.KeyA :
		p.P.X -= p.MoveSpeed * dt
	case ebiten.KeyS :
		p.P.Y -= p.MoveSpeed * dt
	case ebiten.KeyD :
		p.P.X += p.MoveSpeed * dt
	case ebiten.KeyR :
		c.R += gx.Pi * p.ScaleSpeed * dt
	case ebiten.KeyT :
		c.R -= gx.Pi * p.ScaleSpeed * dt
	case ebiten.KeyRightBracket :
		if e.KeyIsPressed(ebiten.KeyShift) {
			p.R -= gx.Pi * 0.3 * dt
		} else {
			p.R += gx.Pi * 0.3 * dt
		}
	case ebiten.KeyF :
		if e.KeyIsPressed(ebiten.KeyShift) {
			c.S.X -= gx.Pi * p.ScaleSpeed * dt
		} else {
			c.S.X += gx.Pi * p.ScaleSpeed * dt
		}
	case ebiten.KeyG :
		if e.KeyIsPressed(ebiten.KeyShift) {
			c.S.Y -= gx.Pi * p.ScaleSpeed * dt
		} else {
			c.S.Y += gx.Pi * p.ScaleSpeed * dt
		}
	case ebiten.KeyZ :
		if e.KeyIsPressed(ebiten.KeyShift) {
			c.RA.X -= gx.Pi * p.MoveSpeed * dt
		} else {
			c.RA.X += gx.Pi * p.MoveSpeed * dt
		}
	case ebiten.KeyX :
		if e.KeyIsPressed(ebiten.KeyShift) {
			c.RA.Y -= gx.Pi * p.MoveSpeed * dt
		} else {
			c.RA.Y += gx.Pi * p.MoveSpeed * dt
		}
	case ebiten.KeyV :
		if e.KeyIsPressed(ebiten.KeyShift) {
			tri.R -= gx.Pi * 0.3 * dt
		} else {
			tri.R += gx.Pi * 0.3 * dt
		}
	case ebiten.KeyLeftBracket :
		if e.KeyIsPressed(ebiten.KeyShift) {
			rect.R -= gx.Pi * 0.3 * dt
		} else {
			rect.R += gx.Pi * 0.3 * dt
		}
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
	
	if rectMove.Vertices().Contained(rect).Len() > 0 ||
			rect.Vertices().Contained(rectMove).Len() > 0 {
		keyStrs = append(keyStrs, "THIS IS SHIT")
	}
	
	e.DebugPrint(i,
		strings.Join(keyStrs, ", "))
	
}

func (d *Debug) IsVisible() bool {return true}

func main() {
	e := gx.NewEngine(&gx.WindowConfig{
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


	text := &gx.Text{}
	text.Text = "Hello, Text"
	text.Color = gx.Color{
		gx.MaxColorV,
		gx.MaxColorV,
		gx.MaxColorV,
		gx.MaxColorV,
	}
	text.Font, err = gx.LoadFont("media/font.ttf", 16)
	if err != nil {
		log.Fatal(err)
	}
	text.Visible = true
	text.Transform = gx.T()
	text.Transform.P = gx.V(100, 100)
	text.Transform.R = gx.Pi/4
	text.Transform.S = gx.V(4, 1)
	
	player = NewPlayer()
	rect = NewRect()
	tri = NewTri()
	
	e.Add(1, &Debug{})
	e.Add(0, player)
	e.Add(-1, rect)
	e.Add(100, tri)
	e.Add(1001, text)
	
	e.Run()
}
