package main

import "fmt"

func main() {
	fmt.Println("=== Control Flow ===")
	fmt.Println("FizzBuzz 1..15:", FirstN(15))
	for _, year := range []int{1900, 2000, 2024, 2025} {
		fmt.Printf("%d leap? %v\n", year, IsLeapYear(year))
	}
	for _, score := range []int{95, 82, 71, 64, 40, 200} {
		fmt.Printf("score=%d grade=%s\n", score, Grade(score))
	}
}
