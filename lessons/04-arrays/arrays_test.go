package main

import "testing"

func TestSum(t *testing.T) {
	if got := Sum([5]int{10, 20, 30, 40, 50}); got != 150 {
		t.Fatalf("Sum = %d, want 150", got)
	}
	if got := Sum([5]int{}); got != 0 {
		t.Fatalf("Sum of zero-value array = %d, want 0", got)
	}
}

func TestMinMax(t *testing.T) {
	minVal, maxVal := MinMax([5]int{3, -1, 9, 4, 0})
	if minVal != -1 || maxVal != 9 {
		t.Fatalf("MinMax = (%d, %d), want (-1, 9)", minVal, maxVal)
	}
}

func TestReverse(t *testing.T) {
	in := [5]int{1, 2, 3, 4, 5}
	out := Reverse(in)
	want := [5]int{5, 4, 3, 2, 1}
	if out != want {
		t.Fatalf("Reverse = %v, want %v", out, want)
	}
	if in != ([5]int{1, 2, 3, 4, 5}) {
		t.Fatalf("Reverse mutated caller: %v", in)
	}
}
