// Package main is lesson 07: functions — first-class, multiple return, variadic, closures.
//
//	go run ./lessons/07-functions
//	go test ./lessons/07-functions
package main

import "errors"

// Add returns a + b. The classic two-arg function.
func Add(a, b int) int { return a + b }

// DivMod returns the quotient and remainder of a / b.
// It returns an error when b == 0 so callers can't silently divide by zero.
func DivMod(a, b int) (quot, rem int, err error) {
	if b == 0 {
		return 0, 0, errors.New("division by zero")
	}
	return a / b, a % b, nil
}

// Factorial returns n! using recursion. Negative input returns 0.
func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// SumAll uses a variadic parameter to accept any number of ints.
func SumAll(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Counter returns a closure that yields 1, 2, 3, ... on each call.
// It's the classic demo for how Go closures capture variables by reference.
func Counter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}
