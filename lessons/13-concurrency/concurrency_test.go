package main

import (
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
)

func TestParallelSquare(t *testing.T) {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
	for _, workers := range []int{1, 2, 4, 8, 32} {
		got := ParallelSquare(in, workers)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("workers=%d: got %v, want %v", workers, got, want)
		}
	}
}

func TestParallelSquare_Empty(t *testing.T) {
	got := ParallelSquare(nil, 4)
	if got == nil || len(got) != 0 {
		t.Fatalf("ParallelSquare(nil) = %v, want non-nil empty", got)
	}
}

func TestSafeCounter(t *testing.T) {
	var c SafeCounter
	var wg sync.WaitGroup
	const N = 10_000
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); c.Inc() }()
	}
	wg.Wait()
	if c.Value() != N {
		t.Fatalf("SafeCounter = %d, want %d", c.Value(), N)
	}
}

func TestAtomicAdd(t *testing.T) {
	var v atomic.Int64
	got := AtomicAdd(&v, 5_000)
	if got != 5_000 {
		t.Fatalf("AtomicAdd = %d, want 5000", got)
	}
}
