package gx

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/surdeus/godat/src/sparsex"
	"github.com/surdeus/godat/src/poolx"
	//"fmt"
	"time"
	"math"
)

// The type is used in all Engine interactions
// where you need floating values.
type Float = float64

// The type represents order of drawing.
type Layer int

// Window configuration type.
type WindowConfig struct {
	Title string
	
	Width,
	Height int
	
	FixedSize,
	Fullscreen,
	VSync bool
}

// The main structure that represents current state of [game] engine.
type Engine struct {
	wcfg *WindowConfig
	layers *sparsex.Sparse[Layer, *poolx.Pool[Drawer]]
	updaters *poolx.Pool[Updater]
	lastTime time.Time
	dt Float
	camera *Camera
	keys []Key
}

type engine Engine

var (
	Infinity = math.MaxFloat64
)

// Return current camera.
func (e *Engine) Camera() *Camera {
	return e.camera
}

// Set new current camera.
func (e *Engine) SetCamera(c *Camera) {
	e.camera = c
}

// Get currently pressed keys.
func (e *Engine) Keys() []Key {
	return e.keys
}

// Returns new empty Engine.
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
			T: Transform{
					S: Vector{1, 1},
			},
		},
		updaters: poolx.New[Updater](),
	}
}

// Add new object considering what
// interfaces it implements.
func (e *Engine) Add(l Layer, b any) {
	starter, ok := b.(Starter)
	if ok {
		starter.Start(e)
	}
	
	updater, ok := b.(Updater)
	if ok {
		e.addUpdater(updater)
	}

	drawer, ok := b.(Drawer)
	if ok {
		e.addDrawer(l, drawer)
	}
}

// Delete object from Engine.
func (e *Engine) Del(b any, v ...any) {
	deleter, ok := b.(Deleter)
	if ok {
		deleter.Delete(e, v...)
	}
	
	drawer, ok := b.(Drawer)
	if ok {
		for layer := range e.layers.Vals() {
			layer.V.Del(drawer)
		}
	}
	
	updater, ok := b.(Updater)
	if ok {
		e.updaters.Del(updater)
	}
	
}

func (e *Engine) addDrawer(l Layer, d Drawer) {
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

func (e *Engine) addUpdater(b Updater) {
	e.updaters.Append(b)
}

func (e *engine) Update() error {
	var err error
	eng := (*Engine)(e)

	e.keys = inpututil.
		AppendPressedKeys(e.keys[:0])

	e.dt = time.Since(e.lastTime).Seconds()
	for p := range eng.updaters.Range() {
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
			if !pj.V.IsVisible() {
				continue
			}
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
	
	ebiten.SetVsyncEnabled(e.wcfg.VSync)

	e.lastTime = time.Now()
	return ebiten.RunGame((*engine)(e))
}

