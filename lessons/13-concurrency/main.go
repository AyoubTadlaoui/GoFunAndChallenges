package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("=== Concurrency ===")
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("input    :", in)
	fmt.Println("squared  :", ParallelSquare(in, 4))

	var c SafeCounter
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); c.Inc() }()
	}
	wg.Wait()
	fmt.Println("mutex   counter:", c.Value())

	var a atomic.Int64
	fmt.Println("atomic  counter:", AtomicAdd(&a, 1000))
}
