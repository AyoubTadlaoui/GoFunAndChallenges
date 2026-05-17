// Package main is lesson 11: idiomatic error handling, sentinels, and wrapping.
//
//	go run ./lessons/11-errors
//	go test ./lessons/11-errors
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

// ErrDivByZero is a sentinel error — callers can check for it with errors.Is.
var ErrDivByZero = errors.New("division by zero")

// Divide returns a/b, or ErrDivByZero when b == 0.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivByZero
	}
	return a / b, nil
}

// ReadFile opens path, reads it fully, and returns the contents.
// Wraps the underlying error with %w so callers can use errors.Is / errors.As.
func ReadFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("ReadFile open %q: %w", path, err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("ReadFile read %q: %w", path, err)
	}
	return string(data), nil
}

// Three-layer error wrapping demo. The previous version of this code had a
// subtle bug — `err := layer1` captured the function value (not its result),
// so `err != nil` was always true and the comparison was meaningless. Here
// each layer actually calls the next and wraps the returned error.

func layer1() error {
	return errors.New("disk is on fire")
}

func layer2() error {
	if err := layer1(); err != nil {
		return fmt.Errorf("layer2: %w", err)
	}
	return nil
}

// Layer3 is exported so tests and main can drive it.
func Layer3() error {
	if err := layer2(); err != nil {
		return fmt.Errorf("layer3: %w", err)
	}
	return nil
}
