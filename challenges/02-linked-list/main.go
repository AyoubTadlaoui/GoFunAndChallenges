package main

import "fmt"

func main() {
	l := New[int]()
	l.PushBack(2)
	l.PushBack(3)
	mid := l.PushBack(4)
	l.PushBack(5)
	l.PushFront(1)

	fmt.Println("=== Doubly Linked List ===")
	fmt.Println("list  :", l)
	fmt.Println("len   :", l.Len())

	l.Remove(mid)
	fmt.Println("after remove 4 :", l)

	l.Reverse()
	fmt.Println("after reverse  :", l)
}
