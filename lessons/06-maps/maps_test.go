package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestWordCount(t *testing.T) {
	got := WordCount(strings.Fields("go go fast slow go"))
	want := map[string]int{"go": 3, "fast": 1, "slow": 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("WordCount = %v, want %v", got, want)
	}
}

func TestTopN(t *testing.T) {
	counts := map[string]int{"go": 3, "rust": 2, "zig": 2, "c": 1}
	got := TopN(counts, 3)
	want := []string{"go", "rust", "zig"} // ties broken alphabetically
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("TopN = %v, want %v", got, want)
	}
	if got := TopN(counts, 99); len(got) != 4 {
		t.Fatalf("TopN over-asked = %v, want all 4", got)
	}
}

func TestMergeCounts(t *testing.T) {
	a := map[string]int{"x": 1, "y": 2}
	b := map[string]int{"y": 3, "z": 4}
	got := MergeCounts(a, b)
	want := map[string]int{"x": 1, "y": 5, "z": 4}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("MergeCounts = %v, want %v", got, want)
	}
	// inputs untouched
	if a["y"] != 2 || b["y"] != 3 {
		t.Fatalf("MergeCounts mutated inputs: a=%v b=%v", a, b)
	}
}
