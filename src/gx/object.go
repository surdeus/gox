package gx

type Behaver interface {
	Start()
	Update()
	GetObject() *Object
}

// The object type represents
// basic information for interaction
// with the engine.
type Object struct {
	T Transform
}

func (o *Object) Start() {}
func (o *Object) Update() {}
func (o *Object) GetObject() *Object {
	return o
}

