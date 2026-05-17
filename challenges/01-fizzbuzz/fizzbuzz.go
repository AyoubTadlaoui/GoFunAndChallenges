// Package main is challenge 01: classic FizzBuzz with an io.Writer twist.
//
//	go run ./challenges/01-fizzbuzz
//	go test ./challenges/01-fizzbuzz
//
// The twist: writing to io.Writer instead of stdout makes the function testable.
package main

import (
	"fmt"
	"io"
)

// FizzBuzz prints 1..n to w. Multiples of 3 become "Fizz", of 5 become "Buzz",
// of 15 become "FizzBuzz". Each entry is followed by a newline.
func FizzBuzz(w io.Writer, n int) error {
	for i := 1; i <= n; i++ {
		var line string
		switch {
		case i%15 == 0:
			line = "FizzBuzz"
		case i%3 == 0:
			line = "Fizz"
		case i%5 == 0:
			line = "Buzz"
		default:
			line = fmt.Sprintf("%d", i)
		}
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return nil
}
