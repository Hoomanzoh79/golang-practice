package main

import "math"

// Abser is an interface that defines a single behavior.
// In Go, interfaces specify WHAT a type can do, not WHAT it is.
//
// Any type that has a method with the exact signature:
//     Abs() float64
// automatically implements this interface.
// There is no "implements" keyword in Go.
type Abser interface {
	Abs() float64
}

// MyFloat is a custom type based on float64.
// Defining methods on basic types (via type aliasing) is common in Go.
type MyFloat float64

// Abs implements the Abser interface for MyFloat.
//
// This uses a VALUE receiver, which means:
// - The method works on a copy of the value
// - Both MyFloat and *MyFloat satisfy the interface
//
// Value receivers are typically used for small, immutable types.
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Point is a regular struct type.
type Point struct {
	X, Y float64
}

// Abs implements the Abser interface for *Point.
//
// This method has a POINTER receiver, which means:
// - ONLY *Point implements Abser (not Point)
// - The method can access the original value without copying
//
// Using pointer receivers is common for structs.
func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}
