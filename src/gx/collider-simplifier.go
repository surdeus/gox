package gx

// Implementing the interface lets 
// the engine to work faster about
// collisions because it first checks
// if the the bigger rectangles that
// contain more complicated structure
// do collide.
type ColliderSimplifier interface {
	ColliderSimplify() Rect
}

