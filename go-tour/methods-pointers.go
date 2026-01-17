package main

import (
	"math"
)

type Rectangle struct {
	X, Y float64
}

// Use a value receiver when:
// The method only reads data, not modifies it.
// The type is small and cheap to copy.
// You are certain the receiver will never need to be nil.
// You don’t need to update internal state.
func (r Rectangle) Abs() float64 {
	return math.Sqrt(r.X*r.X + r.Y*r.Y)
}

// Use a pointer receiver when:
// You want to modify the receiver’s fields inside the method.
// The method must handle a nil receiver safely.
// Other methods on the type already use pointer receivers — for consistency.
// The type has large data, and you want to avoid expensive copying.
func (r *Rectangle) Scale(f float64) {
	r.X = r.X * f
	r.Y = r.Y * f
}
