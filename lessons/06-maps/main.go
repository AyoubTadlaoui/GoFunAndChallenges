package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go is fun go is fast go is simple rust is fast zig is small"
	words := strings.Fields(text)
	counts := WordCount(words)

	fmt.Println("=== Maps ===")
	fmt.Println("text   :", text)
	fmt.Println("counts :", counts)
	fmt.Println("top 3  :", TopN(counts, 3))

	more := WordCount(strings.Fields("go is awesome go is great"))
	fmt.Println("merged :", MergeCounts(counts, more))
}
