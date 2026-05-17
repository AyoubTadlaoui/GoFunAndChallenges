package main

import (
	"math"
	"testing"
)

const epsilon = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

func TestRectangleArea(t *testing.T) {
	cases := []struct {
		name          string
		length, width float64
		want          float64
	}{
		{"unit square", 1, 1, 1},
		{"3x4", 3, 4, 12},
		{"zero width", 5, 0, 0},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := RectangleArea(tc.length, tc.width); !almostEqual(got, tc.want) {
				t.Fatalf("RectangleArea(%v, %v) = %v, want %v", tc.length, tc.width, got, tc.want)
			}
		})
	}
}

func TestCircleArea(t *testing.T) {
	if got := CircleArea(1); !almostEqual(got, math.Pi) {
		t.Fatalf("CircleArea(1) = %v, want %v", got, math.Pi)
	}
	if got := CircleArea(0); !almostEqual(got, 0) {
		t.Fatalf("CircleArea(0) = %v, want 0", got)
	}
}

func TestCircleCircumference(t *testing.T) {
	if got := CircleCircumference(1); !almostEqual(got, 2*math.Pi) {
		t.Fatalf("CircleCircumference(1) = %v, want %v", got, 2*math.Pi)
	}
}

func TestDescribeType(t *testing.T) {
	cases := []struct {
		name string
		in   any
		want string
	}{
		{"int", 42, "integer"},
		{"int64", int64(42), "integer"},
		{"float64", 3.14, "float"},
		{"string", "hello", "string"},
		{"bool", true, "boolean"},
		{"slice falls through", []int{1, 2}, "unknown"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := DescribeType(tc.in); got != tc.want {
				t.Fatalf("DescribeType(%v) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
