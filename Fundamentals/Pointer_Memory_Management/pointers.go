package pointers_and_memo

import (
	"fmt"
)

func Pointers() {
	fmt.Println("-------------------- Pointers and Memory Managemnt -----------------------")
	fmt.Println(" - Create a program that demonstrates the use of pointers to swap the values of two variables.")

	x := "Before "
	y := "Switch"
	z := "After "
	var val1, val2 int
	val1 = 5
	val2 = 2024
	fmt.Println(x+y, "\nvalue 1: ", val1, "\nvalue 2: ", val2)
	swap := func(a, b *string, v1, v2 *int) {
		temp := *a
		*a = *b
		*b = temp

		temp2 := *v1
		*v1 = *v2
		*v2 = temp2
	}

	swap(&x, &z, &val1, &val2)
	fmt.Println(x+y, "\nvalue 1: ", val1, "\nvalue 2: ", val2)

	fmt.Println("- Write a function that accepts a pointer to an integer and increments its value.")
	x2 := 0
	incrementPointer := func(a *int) {
		*a++
	}
	incrementPointer(&x2)
	fmt.Println(x2)
}
