package gx

// The type is used in all Engine interactions
// where you need floating values.
type Float = float64

const (
	MaxFloat = math.MaxFloat64
)

// Returns square of the value.
func Sqr(v Float) Float {
	return v * v
}

