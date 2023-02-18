package gx

type Behaver interface {
	Start(*Engine)
	Update(*Engine) error
}

// The object type represents
// basic information for interaction
// with the engine.
type Object struct {
	T Transform
}

func (o *Object) Start(e *Engine) {}
func (o *Object) Update(e *Engine) error {
	return nil
}
func (o *Object) GetObject() *Object {
	return o
}

