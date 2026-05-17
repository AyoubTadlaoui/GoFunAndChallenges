package main

import "fmt"

func main() {
	lib := New()
	_ = lib.Add(Book{ID: "1", Title: "The Go Programming Language", Author: "Donovan & Kernighan", Year: 2015})
	_ = lib.Add(Book{ID: "2", Title: "100 Go Mistakes and How to Avoid Them", Author: "Teiva Harsanyi", Year: 2022})
	_ = lib.Add(Book{ID: "3", Title: "Learning Go", Author: "Jon Bodner", Year: 2024})

	fmt.Println("=== Library ===")
	fmt.Println("Available:", lib.AvailableTitles())

	if err := lib.Lend("1", "Atlas"); err != nil {
		fmt.Println("lend err:", err)
	}
	if err := lib.Lend("1", "Mehdi"); err != nil {
		fmt.Println("expected double-lend err:", err)
	}

	if b, err := lib.Get("1"); err == nil {
		fmt.Printf("book 1: %q borrowed by %q\n", b.Title, b.Borrower)
	}

	_ = lib.Return("1")
	fmt.Println("After return:", lib.AvailableTitles())
}
