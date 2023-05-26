package gx

// Implementing the interface type
// will call the function OnStart
// when first appear on scene BEFORE
// the OnUpdate.
// The v value will be get from Add function.
type Starter interface {
	Start(*Engine, ...any)
}

// Implementing the interface type
// will call the function on each
// engine iteration.
type Updater interface {
	Update(*Engine) error
}

// Implementing the interface type
// will call the function on deleting
// the object.
type Deleter interface {
	Delete(*Engine, ...any)
}

