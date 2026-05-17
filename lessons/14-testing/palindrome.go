// Package main is lesson 14: testing — table-driven tests, benchmarks, fuzzing.
//
//	go test ./lessons/14-testing -v
//	go test ./lessons/14-testing -bench=.
//	go test ./lessons/14-testing -fuzz=FuzzPalindrome -fuzztime=5s
package main

import (
	"strings"
	"unicode"
)

// IsPalindrome reports whether s reads the same forwards and backwards,
// ignoring case and non-letter/digit runes.
func IsPalindrome(s string) bool {
	clean := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			clean = append(clean, unicode.ToLower(r))
		}
	}
	for i, j := 0, len(clean)-1; i < j; i, j = i+1, j-1 {
		if clean[i] != clean[j] {
			return false
		}
	}
	return true
}

// IsPalindromeSimple is a deliberately less efficient version used to show how
// to benchmark two implementations against each other.
func IsPalindromeSimple(s string) bool {
	s = strings.ToLower(s)
	rev := []rune{}
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			rev = append([]rune{r}, rev...) // O(n²) on purpose
		}
	}
	var clean []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			clean = append(clean, r)
		}
	}
	return string(rev) == string(clean)
}
