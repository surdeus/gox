package gx

// Implementing the interface type
// will call the function OnStart
// when first appear on scene BEFORE
// the OnUpdate.
type Starter interface {
	Start(*Engine)
}

// Implementing the interface type
// will call the function on each
// engine iteration.
type Updater interface {
	Update(*Engine) error
}

// The general interface for 
type Behaver interface {
	Starter
	Updater
}

