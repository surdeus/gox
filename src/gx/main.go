package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/surdeus/godat/src/sparsex"
	"github.com/surdeus/godat/src/poolx"
	//"fmt"
	"time"
)

// The type represents order of drawing.
type Layer int

type WindowConfig struct {
	Title string
	Width, Height int
	FixedSize bool
}

type Engine struct {
	wcfg *WindowConfig
	layers *sparsex.Sparse[Layer, *poolx.Pool[Drawer]]
	behavers *poolx.Pool[Behaver]
	lastTime time.Time
	dt Float
	camera *Camera
	keys []Key
}

type engine Engine

func (e *Engine) Camera() *Camera {
	return e.camera
}

func (e *Engine) SetCamera(c *Camera) {
	e.camera = c
}

func (e *Engine) Keys() []Key {
	return e.keys
}

func New(
	cfg *WindowConfig,
) *Engine {
	return &Engine{
		wcfg: cfg,
		layers: sparsex.New[
			Layer,
			*poolx.Pool[Drawer],
		](true),
		camera: &Camera{
			Object: &Object{
				T: Transform{
					S: Vector{1, 1},
				},
			},
		},
		behavers: poolx.New[Behaver](),
	}
}

// Add new object considering what
// interfaces it implements.
func (e *Engine) Add(l Layer, b any) {
	beh, ok := b.(Behaver)
	if ok {
		e.AddBehaver(beh)
		beh.Start(e)
	}

	drw, ok := b.(Drawer)
	if ok {
		e.AddDrawer(l, drw)
	}
}

func (e *Engine) AddDrawer(l Layer, d Drawer) {
	g, ok := e.layers.Get(l)
	if !ok {
		layer := poolx.New[Drawer]()
		e.layers.Set(
			l,
			layer,
		)
		layer.Append(d)
	} else {
		g.Append(d)
	}

}

func (e *Engine) AddBehaver(b Behaver) {
	e.behavers.Append(b)
}

func (e *engine) Update() error {
	var err error
	eng := (*Engine)(e)

	e.keys = inpututil.
		AppendPressedKeys(e.keys[:0])

	e.dt = time.Since(e.lastTime).Seconds()
	for p := range eng.behavers.Range() {
		err = p.V.Update(eng)
		if err != nil {
			return err
		}
	}
	e.lastTime = time.Now()

	return nil
}

func (e *engine) Draw(i *ebiten.Image) {
	eng := (*Engine)(e)
	for p := range e.layers.Vals() {
		for pj := range p.V.Range() {
			pj.V.Draw(eng, i)
		}
	}
}

func (e *engine) Layout(ow, oh int) (int, int) {
	if e.wcfg.FixedSize {
		return e.wcfg.Width, e.wcfg.Height
	}

	return ow, oh
}

// Return the delta time duration value.
func (e *Engine) DT() Float {
	return e.dt
}

func (e *Engine) Run() error {
	ebiten.SetWindowTitle(e.wcfg.Title)
	ebiten.SetWindowSize(e.wcfg.Width, e.wcfg.Height)
	ebiten.SetWindowSizeLimits(1, 1, e.wcfg.Width, e.wcfg.Height)

	e.lastTime = time.Now()

	return ebiten.RunGame((*engine)(e))
}

