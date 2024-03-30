package varsConsts

import (
	"fmt"
	"math"
	"reflect"
)

func VarTypes() {
	fmt.Println("------------------- Variables and Constances: ------------------- ")

	/*- Write a program that declares variables of different data types (integers, floats, strings, booleans) and prints their values. */

	fmt.Println("- Write a program that declares variables of different data types (integers, floats, strings, booleans) and prints their values. ")
	fmt.Println()
	var Int int
	var Float float32
	var String string
	var Boolean bool

	Int = 29.0
	Float = 29 / 1.5
	String = " \"This is a String\""
	//Nested Function
	isString := func(v interface{}) bool {
		return reflect.TypeOf(v).Kind() == reflect.String
	}

	Boolean = isString(String)

	fmt.Println("Variable int is an integer datatype with value of:", Int,
		"\nVariable float is a float datatype with a value of:", Float,
		"\nVariable String is a string datatype with a value of:", String,
		"\nVariable Boolean is a boolean datatype with a value of:", Boolean)

}

//    - Create a program that calculates the area of a rectangle using variables to store the length and width.

func RectangleAria() {
	fmt.Println("- Create a program that calculates the area of a rectangle using variables to store the length and width.")
	fmt.Println()
	// Declaring Vars.
	var length, width, aria int
	// Entering the length
	fmt.Print("Enter the Length of the rectangle:")
	fmt.Scan(&length)
	// Entering the width
	fmt.Print("Enter the width of the rectangle:")
	fmt.Scan(&width)
	//Calculate the RectangleAria
	aria = length * width
	fmt.Println("The aria of a rectangle is the product of its length with its width\nThe aria of this Rectangle is:", length, "x", width, "=", aria)
}

//    - Experiment with constants by defining mathematical constants like pi and using them in calculations.

func CircleAriaAndSurface() {
	fmt.Println("- Create a progrm that calculates the Aria and Surface of a Circle")
	fmt.Println()
	//Declaring vars
	var choices int
	var radius, aria, surface float32
	const pi = math.Phi

	//Default
	radius = 5
	//User input prompt
	fmt.Println("Choose what do you want to calculate:")
	//Choosing
	fmt.Println("- Aria of a Circle press 1\n- Surface of a Circle press 2\n- Calculate both the Aria and the Surface of a Circle press 3")
	fmt.Scan(&choices)
	//Radius
	fmt.Print("- Enter the Radius of the Circle:")
	fmt.Scan(&radius)
	//Claculations
	aria = radius * radius * pi
	surface = radius * pi * 2
	// Output of Aria calculation
	ariaPrint := fmt.Sprintln("The Aria of a Circle is pi times radius power 2:", aria)
	//Output of Surface Calculation
	surfacePrint := fmt.Sprintln("The Surface of a Circle is pi times radius times 2:", surface)

	//Default
	choices = 3
	//Conditional result
	if choices == 1 {
		fmt.Println()
		fmt.Println(ariaPrint)

	} else if choices == 2 {
		fmt.Println()
		fmt.Println(surfacePrint)

	} else if choices == 3 {
		fmt.Println()
		fmt.Println(ariaPrint)
		fmt.Println(surfacePrint)
	}
}
