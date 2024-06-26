package shapes

import "math"

//- Function Parameters and Return Values: Passing arguments to functions and returning values.

// Define Circle struct
type Circle struct {
	Radius float64
}

// Define method to calculate surface area of the circle
func (c Circle) Surface() float64 {
	return 2 * math.Pi * c.Radius
}
