package gx

type Behaver interface {
	Start(*Engine)
	Update(*Engine)
}

// The object type represents
// basic information for interaction
// with the engine.
type Object struct {
	T Transform
}

func (o *Object) Start(e *Engine) {}
func (o *Object) Update(e *Engine) {}
func (o *Object) GetObject() *Object {
	return o
}

