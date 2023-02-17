package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/surdeus/godat/src/sparsex"
	"fmt"
)

// The type represents order of drawing.
type Layer int

type WindowConfig struct {
	Title string
	Width, Height int
}

type Engine struct {
	wcfg *WindowConfig
	layers *sparsex.Sparse[Layer, *[]Behaver]
	behavers []Behaver
}

type engine Engine

func New(
	cfg *WindowConfig,
) *Engine {
	return &Engine{
		wcfg: cfg,
		layers: sparsex.New[
			Layer,
			*[]Behaver,
		](true),
	}
}

func (e *Engine) Add(l Layer, b Behaver) {
	g, ok := e.layers.Get(l)
	if !ok {
		e.layers.Set(
			l,
			&[]Behaver{b},
		)
	} else {
		set := append(*g, b)
		*g = set
	}

	e.behavers = append(e.behavers, b)
}

func (e *engine) Update() error {
	eng := (*Engine)(e)
	for _, v := range eng.behavers {
		v.Update(eng)
		fmt.Println(v)
	}

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

