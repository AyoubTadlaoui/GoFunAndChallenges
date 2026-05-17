// Package calc shows what a small utility package looks like. The package
// name (calc) and the import path
// (github.com/AyoubTadlaoui/GoFunAndChallenges/lessons/12-packages/calc)
// are independent things — code uses the package name, the module uses the path.
package calc

import "errors"

// ErrDivByZero is returned when attempting to divide by zero.
var ErrDivByZero = errors.New("calc: divide by zero")

// Add returns a + b.
func Add(a, b float64) float64 { return a + b }

// Sub returns a - b.
func Sub(a, b float64) float64 { return a - b }

// Mul returns a * b.
func Mul(a, b float64) float64 { return a * b }

// Div returns a / b, or ErrDivByZero when b == 0.
func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}
