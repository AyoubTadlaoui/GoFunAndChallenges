package FunctionsAndMethods

import (
	"LearningGO/Fundamentals/Functions_and_Methods/shapes"
	"fmt"
)

//- Function Declaration: Syntax for declaring and calling functions in Go
// - Write a function that takes two integers as input parameters and returns their sum.

func Addition() {

	var a, b float32
	fmt.Println("- Write a function that takes two integers as input parameters and returns their sum.")
	fmt.Println("Input the first number you wanna add:")
	fmt.Scan(&a)
	fmt.Println("Input the second number you wanna add:")
	fmt.Scan(&b)

	addFunc := func(i, n float32) {
		fmt.Println("- Write a function that takes two integers as input parameters and returns their sum.")
		fmt.Println()
		var result = i + n
		fmt.Println(i, "+", n, "=", result)
	}

	addFunc(a, b)

}

// - Implement a function to calculate the factorial of a given integer using recursion.
func Factorial() {
	//declare the variable
	var i int
	var calcFactorial func(int) int
	//Prompt
	fmt.Println("Enter the number you want to calculate its factorial.")
	fmt.Scan(&i)

	//function to calc factorial
	calcFactorial = func(n int) int {
		if n == 0 || n == 1 {
			return 1
		}
		return n * calcFactorial(n-1)
	}
	//Result
	var result = calcFactorial(i)
	//Print result:
	fmt.Println(i, "! =", result)
}

// - Define a method on a struct representing a geometric shape to calculate its area.

func MethodStruct() {

	fmt.Println("---------------- Methods and Functions ------------------")
	fmt.Println(" - Define a method on a struct representing a geometric shape to calculate its area.")
	fmt.Println()
	//===========Building the  Structures=============

	type Rectangle struct {
		length, width float64
	}

	// ============ Build Interfaces =================
	// Define Shape interface

	// type Shape interface {
	// 	Aria() float64
	// }

	//============== Rectangle methods ===============
	//Let's make a method to calculate the Aria of thie Structure(i.e assosition)
	recAria := func(r Rectangle) float64 {
		return r.length * r.width
	}
	//Prompt to input width and length:
	var w, l float64
	fmt.Println("Enter the width of the rectangle:")
	fmt.Scan(&w)
	fmt.Println("Enter the length of the rectangle:")
	fmt.Scan(&l)
	//Creating an instance:
	recData := Rectangle{width: w, length: l}
	//printing the result of aria
	fmt.Println("The aria of this rectengle is:", w, "x", l, "=", recAria(recData))

	//============ Circle Methods =================

	var r float64
	fmt.Println("Input the Radius of the Circle!")
	fmt.Scan(&r)

	cInData := shapes.Circle{Radius: r}
	result := cInData.Surface()
	fmt.Println("The Aria of this circle is:", result)

}

func Veriadic() {

	summation := func(numbers ...int) (int, int) {
		total := 0
		totalIndices := 0
		for index, sum := range numbers {
			totalIndices += index
			total += sum
		}

		return total, totalIndices
	}
	slice1 := []int{1, 2, 3}
	sum1, sumindices1 := summation(slice1...)
	println(sum1, sumindices1)
}
