// Package main is lesson 01: the first runnable Go program.
//
// Run it with:
//
//	go run ./lessons/01-hello
package main

import "fmt"

func main() {
	fmt.Println(Greet("Gopher"))
}

// Greet returns a friendly welcome line for name.
func Greet(name string) string {
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("Hello, %s! Welcome to Go.", name)
}
