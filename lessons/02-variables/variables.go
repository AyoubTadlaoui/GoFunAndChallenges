// Package main is lesson 02: variables, constants, and small math helpers.
//
// Run it with:
//
//	go run ./lessons/02-variables
//
// Tests live next to the code:
//
//	go test ./lessons/02-variables
package main

import "math"

// Pi is the mathematical constant π. The Go standard library exposes it as
// math.Pi; we re-export it as a named constant here so the lesson reads cleanly.
const Pi = math.Pi

// RectangleArea returns length * width. Real-world measurements are float64.
func RectangleArea(length, width float64) float64 {
	return length * width
}

// CircleArea returns π·r² — the area enclosed by a circle of radius r.
func CircleArea(radius float64) float64 {
	return Pi * radius * radius
}

// CircleCircumference returns 2·π·r — the distance around the circle.
//
// Earlier versions of this course called this "surface", which is wrong:
// the 2D surface of a disk is its area; 2πr is the perimeter.
func CircleCircumference(radius float64) float64 {
	return 2 * Pi * radius
}

// DescribeType returns a short label for the kind of a value, using a type
// switch over `any`.
func DescribeType(v any) string {
	switch v.(type) {
	case int, int32, int64:
		return "integer"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case bool:
		return "boolean"
	default:
		return "unknown"
	}
}
