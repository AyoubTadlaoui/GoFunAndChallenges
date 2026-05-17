package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3, 4}); got != 10 {
		t.Fatalf("Sum = %d, want 10", got)
	}
	if got := Sum(nil); got != 0 {
		t.Fatalf("Sum(nil) = %d, want 0", got)
	}
}

func TestFilter(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6}
	got := Filter(in, func(n int) bool { return n%2 == 0 })
	want := []int{2, 4, 6}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Filter = %v, want %v", got, want)
	}
	// caller is not mutated
	if !reflect.DeepEqual(in, []int{1, 2, 3, 4, 5, 6}) {
		t.Fatalf("Filter mutated input: %v", in)
	}
}

func TestMap(t *testing.T) {
	got := Map([]int{1, 2, 3}, func(n int) string {
		return [...]string{"one", "two", "three"}[n-1]
	})
	want := []string{"one", "two", "three"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Map = %v, want %v", got, want)
	}
}

func TestUnique(t *testing.T) {
	got := Unique([]string{"a", "b", "a", "c", "b", "a"})
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Unique = %v, want %v", got, want)
	}
	if got := Unique([]int{}); len(got) != 0 {
		t.Fatalf("Unique([]) = %v, want empty", got)
	}
}
