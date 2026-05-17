// Package main is lesson 06: maps — Go's hash table.
//
//	go run ./lessons/06-maps
//	go test ./lessons/06-maps
package main

import "sort"

// WordCount returns how many times each word appears in words.
func WordCount(words []string) map[string]int {
	counts := make(map[string]int, len(words))
	for _, w := range words {
		counts[w]++
	}
	return counts
}

// TopN returns the n words with the highest counts, ties broken alphabetically.
// If n is larger than the number of distinct words, it returns all of them.
func TopN(counts map[string]int, n int) []string {
	words := make([]string, 0, len(counts))
	for w := range counts {
		words = append(words, w)
	}
	sort.Slice(words, func(i, j int) bool {
		if counts[words[i]] != counts[words[j]] {
			return counts[words[i]] > counts[words[j]]
		}
		return words[i] < words[j]
	})
	if n > len(words) {
		n = len(words)
	}
	return words[:n]
}

// MergeCounts returns a new map containing every key from a and b, with values summed.
// Neither input is mutated.
func MergeCounts(a, b map[string]int) map[string]int {
	out := make(map[string]int, len(a)+len(b))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		out[k] += v
	}
	return out
}
