// Package main is lesson 03: if / else, for, switch.
//
//	go run ./lessons/03-control-flow
//	go test ./lessons/03-control-flow
package main

import "strings"

// FizzBuzz classifies n by its divisibility:
//   - "FizzBuzz" if divisible by 15
//   - "Fizz"     if divisible by 3
//   - "Buzz"     if divisible by 5
//   - otherwise the decimal form of n
func FizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		// Build the digits manually to avoid pulling in strconv here — readers
		// see the algorithm before the standard library does the heavy lifting.
		return itoa(n)
	}
}

// IsLeapYear reports whether year is a Gregorian leap year.
//
// Rule: divisible by 4, except century years not divisible by 400.
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	return year%4 == 0
}

// Grade returns the letter grade for a percent score in [0, 100].
// Out-of-range inputs return "?".
func Grade(percent int) string {
	switch {
	case percent < 0 || percent > 100:
		return "?"
	case percent >= 90:
		return "A"
	case percent >= 80:
		return "B"
	case percent >= 70:
		return "C"
	case percent >= 60:
		return "D"
	default:
		return "F"
	}
}

// itoa converts a (possibly negative) int to its decimal string representation
// using only loops and arithmetic. It's a deliberate teaching version — prefer
// strconv.Itoa in real code.
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if neg {
		n = -n
	}
	var digits []byte
	for n > 0 {
		digits = append(digits, byte('0'+n%10))
		n /= 10
	}
	// reverse
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	out := string(digits)
	if neg {
		out = "-" + out
	}
	return out
}

// FirstN returns FizzBuzz for 1..n joined with spaces — handy for the demo.
func FirstN(n int) string {
	parts := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		parts = append(parts, FizzBuzz(i))
	}
	return strings.Join(parts, " ")
}
