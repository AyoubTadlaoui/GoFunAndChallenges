// Package main is lesson 05: slices — Go's dynamic sequence.
//
//	go run ./lessons/05-slices
//	go test ./lessons/05-slices
package main

// Sum adds every element of s.
func Sum(s []int) int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}

// Filter returns a new slice containing the elements of s for which keep is true.
// It allocates a fresh backing array so callers can't accidentally mutate s.
func Filter[T any](s []T, keep func(T) bool) []T {
	out := make([]T, 0, len(s))
	for _, v := range s {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}

// Map returns a new slice where each element is f(s[i]).
func Map[T, U any](s []T, f func(T) U) []U {
	out := make([]U, len(s))
	for i, v := range s {
		out[i] = f(v)
	}
	return out
}

// Unique returns the elements of s in original order, with later duplicates dropped.
func Unique[T comparable](s []T) []T {
	seen := make(map[T]struct{}, len(s))
	out := make([]T, 0, len(s))
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		out = append(out, v)
	}
	return out
}
