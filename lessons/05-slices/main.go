package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("=== Slices ===")
	fmt.Println("nums       =", nums)
	fmt.Println("sum        =", Sum(nums))

	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("evens      =", evens)

	squared := Map(nums, func(n int) int { return n * n })
	fmt.Println("squared    =", squared)

	dups := []string{"go", "rust", "go", "zig", "rust", "go"}
	fmt.Println("unique     =", Unique(dups))
}
