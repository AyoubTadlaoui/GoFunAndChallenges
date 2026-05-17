package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== Errors ===")

	if q, err := Divide(10, 4); err == nil {
		fmt.Printf("10 / 4 = %.2f\n", q)
	}

	if _, err := Divide(1, 0); err != nil {
		fmt.Println("1 / 0 ->", err)
		if errors.Is(err, ErrDivByZero) {
			fmt.Println("(matched sentinel ErrDivByZero)")
		}
	}

	if _, err := ReadFile("definitely-not-there.txt"); err != nil {
		fmt.Println("read missing  ->", err)
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("(matched os.ErrNotExist via errors.Is)")
		}
	}

	fmt.Println("Layer3 ->", Layer3())
}
