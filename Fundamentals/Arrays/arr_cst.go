package arrayconst

import (
	"fmt"
	"math"
)

func Arrays() {
	// Print a header for the Arrays section
	fmt.Println("------------------- Arrays --------------------")

	// Define an array of integers and print each element
	var arrayx = [...]int{10, 20, 30}
	arrayx2 := [3]int{10, 20, 30}
	array3 := []int{}                     // Create an empty slice
	slicex3 := append(array3, 10, 20, 30) // Append elements to the empty slice
	fmt.Println(arrayx, "\n", arrayx2, "\n", slicex3, "\n", (arrayx == arrayx2))

	// Calculate the sum of all elements in an array
	fmt.Println("- Calculate the sum of all elements in an array.")
	sumarrayx := 0
	// Summation with iteration
	for _, value := range arrayx {
		sumarrayx += value
	}
	fmt.Println("The sum by Iteration of the element of array x is:", sumarrayx)
	// Emptying the variable Sumarrayx for reusage
	sumarrayx = 0
	// Direct summation
	for i := 0; i < len(arrayx); i++ {
		sumarrayx += arrayx[i]
	}
	fmt.Println("The sum of array x by direct calculation is: ", sumarrayx)

	// Find the maximum and minimum values in an array
	fmt.Println("- Find the maximum and minimum values in an array")
	max := math.MinInt16
	min := math.MaxInt16
	// Finding max and min using math library
	for _, value := range arrayx {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}
	fmt.Println("Using \"math\" library:\nmax value:", max, "\nmin value:", min)
	// Finding max and min without math library
	max = arrayx[0]
	min = arrayx[0]
	for _, value := range arrayx {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}
	fmt.Println("Using first index as min and max value\nmax value:", max, "\nmin value:", min)
}
