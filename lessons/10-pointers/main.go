package main

import "fmt"

func main() {
	fmt.Println("=== Pointers ===")
	a, b := 1, 2
	fmt.Println("before swap:", a, b)
	Swap(&a, &b)
	fmt.Println("after swap :", a, b)

	x := 41
	Inc(&x)
	fmt.Println("after inc  :", x)

	fmt.Println("=== Linked List ===")
	list := NewLinkedList()
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Prepend(0)
	fmt.Println("list  :", list)
	fmt.Println("len   :", list.Len())
	fmt.Println("values:", list.Values())
}
