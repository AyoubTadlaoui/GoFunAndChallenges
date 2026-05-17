package main

import "fmt"

func main() {
	fmt.Println("=== Functions ===")
	fmt.Println("Add(2, 3)         =", Add(2, 3))

	q, r, err := DivMod(17, 5)
	fmt.Printf("DivMod(17, 5)     = q=%d r=%d err=%v\n", q, r, err)
	_, _, err = DivMod(1, 0)
	fmt.Printf("DivMod(1, 0)      = err=%v\n", err)

	fmt.Println("Factorial(6)      =", Factorial(6))
	fmt.Println("SumAll(1..5)      =", SumAll(1, 2, 3, 4, 5))

	next := Counter()
	fmt.Println("Counter calls     =", next(), next(), next())
}
