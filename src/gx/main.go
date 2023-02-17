package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type WindowConfig struct {
	Title string
	Width, Height int
}

type Engine struct {
	wcfg *WindowConfig
}

type engine Engine

func New(
	cfg *WindowConfig,
) *Engine {
	return &Engine{
		wcfg: cfg,
	}
}

func (e *engine) Update() error {
	return nil
}

func (e *engine) Draw(s *ebiten.Image) {
}

func (e *engine) Layout(ow, oh int) (int, int) {
	return e.wcfg.Width, e.wcfg.Height
}

func (e *Engine) Run() error {
	ebiten.SetWindowTitle(e.wcfg.Title)
	ebiten.SetWindowSize(e.wcfg.Width, e.wcfg.Height)

	return ebiten.RunGame((*engine)(e))
}

