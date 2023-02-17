package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/surdeus/godat/src/sparsex"
	//"fmt"
	"time"
)

// The type represents order of drawing.
type Layer int

type WindowConfig struct {
	Title string
	Width, Height int
	FullScreen bool
}

type Engine struct {
	wcfg *WindowConfig
	layers *sparsex.Sparse[Layer, *[]Drawer]
	behavers []Behaver
	lastTime time.Time
	dt Float
}

type engine Engine

func New(
	cfg *WindowConfig,
) *Engine {
	return &Engine{
		wcfg: cfg,
		layers: sparsex.New[
			Layer,
			*[]Drawer,
		](true),
	}
}

// Add new object considering what
// interfaces it implements.
func (e *Engine) Add(l Layer, b any) {
	beh, ok := b.(Behaver)
	if ok {
		e.AddBehaver(beh)
	}

	drw, ok := b.(Drawer)
	if ok {
		e.AddDrawer(l, drw)
	}
}

func (e *Engine) AddDrawer(l Layer, d Drawer) {
	g, ok := e.layers.Get(l)
	if !ok {
		e.layers.Set(
			l,
			&[]Drawer{d},
		)
	} else {
		set := append(*g, d)
		*g = set
	}

}

func (e *Engine) AddBehaver(b Behaver) {
	e.behavers = append(e.behavers, b)
}

func (e *engine) Update() error {
	eng := (*Engine)(e)

	e.dt = time.Since(e.lastTime).Seconds()
	for _, v := range eng.behavers {
		v.Update(eng)
		//fmt.Println(v)
	}
	e.lastTime = time.Now()

	return nil
}


func (e *engine) Draw(i *ebiten.Image) {
	eng := (*Engine)(e)
	for p := range e.layers.Vals() {
		for _, d := range *p.V {
			d.Draw(eng, i)
		}
	}
}

func (e *engine) Layout(ow, oh int) (int, int) {
	return e.wcfg.Width, e.wcfg.Height
}

// Return the delta time duration value.
func (e *Engine) DT() Float {
	return e.dt
}

func (e *Engine) Run() error {
	ebiten.SetWindowTitle(e.wcfg.Title)
	ebiten.SetWindowSize(e.wcfg.Width, e.wcfg.Height)

	return ebiten.RunGame((*engine)(e))
}

