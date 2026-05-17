package main

import "fmt"

func main() {
	cases := []string{
		"racecar",
		"A man, a plan, a canal: Panama",
		"hello",
		"",
	}
	fmt.Println("=== Testing & Benchmarks ===")
	for _, s := range cases {
		fmt.Printf("IsPalindrome(%q) = %v\n", s, IsPalindrome(s))
	}
	fmt.Println()
	fmt.Println("Run the tests:    go test ./lessons/14-testing -v")
	fmt.Println("Run benchmarks:   go test ./lessons/14-testing -bench=.")
	fmt.Println("Run fuzz (5s):    go test ./lessons/14-testing -fuzz=FuzzPalindrome -fuzztime=5s")
}
