package main

import "testing"

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add = %d, want 5", got)
	}
}

func TestDivMod(t *testing.T) {
	q, r, err := DivMod(17, 5)
	if err != nil || q != 3 || r != 2 {
		t.Fatalf("DivMod(17,5) = (%d, %d, %v), want (3, 2, nil)", q, r, err)
	}
	if _, _, err := DivMod(7, 0); err == nil {
		t.Fatal("DivMod(7,0) returned nil error, want non-nil")
	}
}

func TestFactorial(t *testing.T) {
	cases := map[int]int{0: 1, 1: 1, 2: 2, 5: 120, 6: 720, -1: 0}
	for n, want := range cases {
		if got := Factorial(n); got != want {
			t.Errorf("Factorial(%d) = %d, want %d", n, got, want)
		}
	}
}

func TestSumAll(t *testing.T) {
	if got := SumAll(); got != 0 {
		t.Fatalf("SumAll() = %d, want 0", got)
	}
	if got := SumAll(1, 2, 3, 4, 5); got != 15 {
		t.Fatalf("SumAll(1..5) = %d, want 15", got)
	}
	// passing a slice with ... spread
	nums := []int{10, 20, 30}
	if got := SumAll(nums...); got != 60 {
		t.Fatalf("SumAll(nums...) = %d, want 60", got)
	}
}

func TestCounter(t *testing.T) {
	c := Counter()
	for i := 1; i <= 3; i++ {
		if got := c(); got != i {
			t.Fatalf("Counter call %d = %d, want %d", i, got, i)
		}
	}
	// independent counters don't share state
	c2 := Counter()
	if got := c2(); got != 1 {
		t.Fatalf("Second Counter call = %d, want 1", got)
	}
}
