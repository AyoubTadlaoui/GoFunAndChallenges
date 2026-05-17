package main

import "fmt"

func main() {
	a := [5]int{10, 20, 30, 40, 50}
	fmt.Println("=== Arrays ===")
	fmt.Println("a       =", a)
	fmt.Println("sum     =", Sum(a))
	minVal, maxVal := MinMax(a)
	fmt.Println("min,max =", minVal, maxVal)
	fmt.Println("reverse =", Reverse(a))
	fmt.Println("a still =", a, "(arrays are value types — Reverse can't mutate the caller)")
}
