// Package main is lesson 13: goroutines, channels, sync primitives.
//
//	go run ./lessons/13-concurrency
//	go test ./lessons/13-concurrency
//
// The mantra is: "Don't communicate by sharing memory; share memory by communicating."
package main

import (
	"sync"
	"sync/atomic"
)

// ParallelSquare squares every element of in concurrently using a worker pool
// of size workers, and returns the results in the original order.
//
// If workers < 1 it falls back to 1. If in is empty, an empty (non-nil) slice
// is returned and no goroutines are spawned.
func ParallelSquare(in []int, workers int) []int {
	out := make([]int, len(in))
	if len(in) == 0 {
		return out
	}
	if workers < 1 {
		workers = 1
	}
	if workers > len(in) {
		workers = len(in)
	}

	type job struct {
		idx int
		val int
	}
	jobs := make(chan job)

	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				out[j.idx] = j.val * j.val
			}
		}()
	}

	for i, v := range in {
		jobs <- job{idx: i, val: v}
	}
	close(jobs)
	wg.Wait()
	return out
}

// SafeCounter is a goroutine-safe counter using sync.Mutex.
// It's the canonical example of guarding shared state with a lock.
type SafeCounter struct {
	mu sync.Mutex
	n  int
}

// Inc increments the counter under the lock.
func (c *SafeCounter) Inc() {
	c.mu.Lock()
	c.n++
	c.mu.Unlock()
}

// Value returns the current count.
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

// AtomicAdd runs n goroutines that each Add(1) to v using sync/atomic, and
// returns the final value. Useful for comparing to mutexes for hot-path counters.
func AtomicAdd(v *atomic.Int64, n int) int64 {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			v.Add(1)
		}()
	}
	wg.Wait()
	return v.Load()
}
