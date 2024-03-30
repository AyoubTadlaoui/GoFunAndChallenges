package testing

import (
	"fmt"
	"reflect"
)

func Pointers() {

	var i int
	z := 10
	r := 100
	pointerToR := &r
	fmt.Println(" r variables that has assigned an int of:", r, " islocated in the pointed adress: ", pointerToR)
	fmt.Println("The type of pointerToR is: ", reflect.TypeOf(pointerToR))
	fmt.Println("i has no value and will be pointed a new value")
	fmt.Scan(&i)
	fmt.Println("z is already assigned an int 10 and will will try to point a new value to it")
	fmt.Scan(&z)

	outPrint := func(a int, b int) string {

		return fmt.Sprintf("%d is the value assigned to i\n%p is the memory adress of i\n%d is the new value assigned to z ", a, &a, b)
	}

	result := outPrint(i, z)
	fmt.Print(result)
}
