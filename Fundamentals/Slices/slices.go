package slices

import "fmt"

func Slices() {
	// Print a header for the Slices section
	fmt.Println("----------------- Slices ---------------------")

	// Tests of Slices
	// - Use `make` to create a slice with a specific length and
	arrayx := [...]int{2016, 2017, 2018, 2019, 2020, 2024}
	slice := make([]int, 3, 5)
	fmt.Println(slice)
	slice = arrayx[0:2]
	fmt.Println(slice)
	slice2 := []int{1, 2, 3}
	slice2 = append(slice2, 4, 5)
	subSlice2 := slice2[4]
	fmt.Println(slice2, "\n", subSlice2)
	// - Create a slice of strings and append new elements to it
	fmt.Println("- Create a slice of strings and append new elements to it")
	stringSlice := make([]string, 4)
	stringSlice[0] = "BOLO"
	stringSlice[1] = "Pinkoko"
	stringSlice[2] = "BOSS"
	stringSlice[3] = "MrStrikeforce"
	// Create a slice of strings and append new elements to it
	fmt.Println(stringSlice[0], stringSlice[1], stringSlice[2], stringSlice[3])

	stringSlice2 := []string{"Friends", "Rich", "Wealthy", "Happy", "Sweet revange"}
	stringSliceAppend := append(stringSlice2, "2024")
	fmt.Println(stringSlice2)
	fmt.Println(stringSliceAppend)
	// - Slice an existing slice to create a subset.
	fmt.Println("- Slice an existing slice to create a subset")
	subSliceValues := []string{stringSlice2[3], stringSlice2[0], stringSlice2[4]}
	fmt.Println(subSliceValues)

	// Printing the elements dynamically using loops
	stringSlices3 := []string{"Money", "Power", "Wealth", "Iron Will", "Discipline"}
	indxSlice3 := []int{0, 2}
	var subSlice3 []string
	for _, indx := range indxSlice3 {
		subSlice3 = append(subSlice3, stringSlices3[indx])
	}
	fmt.Println(subSlice3)

	// Iterate over a slice and print each element
	fmt.Println("- Iterate over a slice and print each element.")
	for _, value := range stringSlices3 {
		fmt.Println(value)
	}
}
