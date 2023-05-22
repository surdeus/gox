package gx

// Implementing the interface lets 
// the engine to work faster about
// collisions because it first checks
// if the the bigger collider that
// contain more complicated structure
// do collide.
type ColliderSimplifier interface {
	ColliderSimplify() Triangle
}

// The structure represents all
// information on collisions.
type Collision struct {
	Current, With any
}

// Implementing the interface lets the engine
// to determine if the object collides with anything.
// Mostly will use the Collide function with some
// inner structure field as first argument.
type Collider interface {
	Collides(Collider) *Collision
}

// happening collision getting the Collision as
// argument.
type CollideEventer interface {
	Collide(*Collision)
}

// Single function for all collision to remove 
// functionality duplicating from the archtecture.
// Returns the collision if there is and nil if there
// is no collision.
/*func Collide(c1, c2 any) bool {
}

func triangleCollidesPoint(t Triangle, p Point) *Collision {
}

func triangleCollidesTriangle(t1, t2 Triangle) *Collision
*/

