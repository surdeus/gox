package gx

// Implementing the interface lets 
// the engine to work faster about
// collisions because it first checks
// if the the bigger collider that
// contain more complicated structure
// do collide.
type ColliderSimplifier interface {
	ColliderSimplify() Rect
}

// The structure represents all
// information on collision.
type Collision struct {
	Other Collider
}

// Every collider has to implement
// collision with every other type of collider
// for optimization. Not good for custom colliders
// but is fast.
type Collider interface {
	Collides(Collider) *Collision
}

// happening collision getting the Collision as
// argument.
type CollideEventer interface {
	Collide(*Collision)
}

