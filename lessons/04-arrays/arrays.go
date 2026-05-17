// Package main is lesson 04: fixed-size arrays.
//
//	go run ./lessons/04-arrays
//	go test ./lessons/04-arrays
//
// In Go, an array's length is part of its type: [3]int and [4]int are
// different types. For dynamic collections, see lesson 05 (Slices).
package main

// Sum returns the sum of all elements in a.
func Sum(a [5]int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}

// MinMax returns the smallest and largest elements of a.
// The array type guarantees there's at least one element.
func MinMax(a [5]int) (minVal, maxVal int) {
	minVal, maxVal = a[0], a[0]
	for _, v := range a[1:] {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	return
}

// Reverse returns a new array with the elements of a in reverse order.
// Arrays are value types in Go — assigning or returning one copies it.
func Reverse(a [5]int) [5]int {
	var out [5]int
	for i, v := range a {
		out[len(a)-1-i] = v
	}
	return out
}
