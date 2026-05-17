// Package main is lesson 08: methods, interfaces, and pointer receivers.
//
//	go run ./lessons/08-methods
//	go test ./lessons/08-methods
package main

import "math"

// Shape is anything with an area. The interface is satisfied implicitly —
// no `implements` keyword.
type Shape interface {
	Area() float64
}

// Rectangle implements Shape.
type Rectangle struct {
	Width, Height float64
}

// Area returns the area of the rectangle. Value receiver — Rectangle is small
// and immutable here, so copying is fine.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle implements Shape.
type Circle struct {
	Radius float64
}

// Area returns πr².
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Counter is a tiny demo of a pointer receiver.
type Counter struct {
	count int
}

// Inc bumps the counter. Pointer receiver — value receivers can't mutate the caller.
func (c *Counter) Inc() {
	c.count++
}

// Count reads the current value. Pointer receiver kept for consistency with Inc;
// methods of the same type should typically use one receiver kind, not both.
func (c *Counter) Count() int {
	return c.count
}

// TotalArea sums the area of every Shape passed in.
// Demonstrates how interfaces flatten heterogeneous collections.
func TotalArea(shapes ...Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}
