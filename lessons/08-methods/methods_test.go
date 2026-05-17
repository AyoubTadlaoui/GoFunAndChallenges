package main

import (
	"math"
	"testing"
)

const epsilon = 1e-9

func TestRectangleArea(t *testing.T) {
	if got := (Rectangle{3, 4}).Area(); got != 12 {
		t.Fatalf("Rectangle area = %v, want 12", got)
	}
}

func TestCircleArea(t *testing.T) {
	got := Circle{Radius: 1}.Area()
	if math.Abs(got-math.Pi) > epsilon {
		t.Fatalf("Circle area = %v, want %v", got, math.Pi)
	}
}

func TestCounter(t *testing.T) {
	var c Counter
	for i := 1; i <= 5; i++ {
		c.Inc()
		if c.Count() != i {
			t.Fatalf("count after Inc #%d = %d, want %d", i, c.Count(), i)
		}
	}
}

func TestTotalArea(t *testing.T) {
	got := TotalArea(Rectangle{2, 3}, Rectangle{1, 1}, Circle{Radius: 1})
	want := 6 + 1 + math.Pi
	if math.Abs(got-want) > epsilon {
		t.Fatalf("TotalArea = %v, want %v", got, want)
	}
}
