package gx

import (
	"math"
)

// The type is used in all Engine interactions
// where you need floating values.
type Float = float64

const (
	MaxFloat = math.MaxFloat64
	Pi = math.Pi
)

// Returns square of the value.
func Sqr(v Float) Float {
	return v * v
}

func RadiansToDegrees(v Float) Float {
	return v
}

func DeegresToRadians(v Float) Float {
	return v
}

