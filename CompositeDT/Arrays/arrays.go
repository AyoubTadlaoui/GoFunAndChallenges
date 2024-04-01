package arrays

import (
	"fmt"
	"math"
)

func Arrays() {
	fmt.Println("------------------- Arrays -----------------")
	fmt.Println("- Define an array of integers and print each element.")

	var arrayx = [...]int{10, 20, 30}
	arrayx2 := [3]int{10, 20, 30}
	array3 := []int{}
	slicex3 := append(array3, 10, 20, 30)
	fmt.Println(arrayx, "\n", arrayx2, "\n", slicex3, "\n", (arrayx == arrayx2))

	fmt.Println("- Calculate the sum of all elements in an array.")
	sumarrayx := 0
	// sumindex := 0
	//Summation with Iteration:
	for _, value := range arrayx {
		sumarrayx += value
	}
	fmt.Println("The sum by Iteration of the element of array x is:", sumarrayx)
	//Emptying the variable Sumarrayx for reusage
	sumarrayx = 0
	//Direct summation
	for i := 0; i < len(arrayx); i++ {
		sumarrayx += arrayx[i]
	}
	fmt.Println("The sum of array x by direct calculation is: ", sumarrayx)
	fmt.Println("- Find the maximum and minimum values in an array")
	max := math.MinInt16
	min := math.MaxInt16

	for _, value := range arrayx {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	fmt.Println("Using \"math\" library:\nmax value:", max, "\nmin value:", min)

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

	fmt.Println("----------------- Slices ---------------------")
	//Tests of Slices
	slice := make([]int, 3, 5)
	fmt.Println(slice)
	slice = arrayx[0:2]
	fmt.Println(slice)
	slice2 := []int{1, 2, 3}
	slice2 = append(slice2, 4, 5)
	subSlice2 := slice2[4]
	fmt.Println(slice2, "\n", subSlice2)
	stringSlice := make([]string, 4)
	stringSlice[0] = "BOLO"
	stringSlice[1] = "Pinkoko"
	stringSlice[2] = "BOSS"
	stringSlice[3] = "MrStrikeforce"
	//- Create a slice of strings and append new elements to it.
	fmt.Println(stringSlice[0], stringSlice[1], stringSlice[2], stringSlice[3])
	stringSlice2 := []string{"Friends", "Rich", "Wealthy", "Happy", "Sweet revange"}
	stringSliceAppend := append(stringSlice2, "2024")
	fmt.Println(stringSlice2)
	fmt.Println(stringSliceAppend)
	subSliceValues := []string{stringSlice2[3], stringSlice2[0], stringSlice2[4]}
	fmt.Println(subSliceValues)
	//Printing the element dynamicaclly using loops
	stringSlices3 := []string{"Money", "Power", "Wealth", "Iron Will", "Discipline"}
	indxSlice3 := []int{0, 2}
	var subSlice3 []string
	for _, indx := range indxSlice3 {
		subSlice3 = append(subSlice3, stringSlices3[indx])
	}
	fmt.Println(subSlice3)
	//	- Iterate over a slice and print each element.

	for _, value := range stringSlices3 {
		fmt.Println(value)
	}

	fmt.Println("------------------------ Maps -----------------------")
	// - Create a map to store the population of different cities.

	mapForDiffCities := map[string]int{
		"Oujda":   494252,
		"Kenitra": 431282,
		"Rabat":   577827,
	}

	// Print the map
	fmt.Println(mapForDiffCities)

	//Append throught it to meet certain conditions

	bigCities := make(map[string]int)

	for cities, population := range mapForDiffCities {

		if population > 450000 {
			bigCities[cities] = population
		}
	}

	fmt.Println("Big Citie(s)", bigCities)

	newCities := map[string]int{

		"Casablanca": 2600000,
	}

	//Merge the two maps
	for city, population := range mapForDiffCities {
		newCities[city] = population
	}
	fmt.Println(newCities)

	// Selecting only the big cities:

	metropolis := make(map[string]int)

	for city, population := range newCities {

		if population > 1000000 {
			metropolis[city] = population
		}
	}
	//Print result
	fmt.Println("Citie(s) with a population bigger then 1 Million: ", metropolis)

	//Print result without the prefix map and []
	for printCity, printPopulation := range metropolis {
		fmt.Printf("Citie(s) with a population bigger then 1 Million: %s with a population of:%d\n", printCity, printPopulation)
	}

}
