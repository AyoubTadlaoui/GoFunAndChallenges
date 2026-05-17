// Lesson 12: packages and modules — how to organize and import.
//
//	go run ./lessons/12-packages
//	go test ./lessons/12-packages/...
package main

import (
	"fmt"
	"strings"

	"github.com/AyoubTadlaoui/GoFunAndChallenges/lessons/12-packages/calc"
)

func main() {
	fmt.Println("=== Packages & Modules ===")
	fmt.Println("calc.Add(2, 3) =", calc.Add(2, 3))
	fmt.Println("calc.Mul(6, 7) =", calc.Mul(6, 7))

	q, err := calc.Div(10, 0)
	fmt.Printf("calc.Div(10,0) = %v, err=%v\n", q, err)

	// Showing that standard-library packages compose the same way.
	fmt.Println("upper          =", strings.ToUpper("learning go"))
}
