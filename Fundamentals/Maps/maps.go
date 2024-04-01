package maps

import "fmt"

func Maps() {
	// Print a header for the Maps section
	fmt.Println("------------------------ Maps -----------------------")

	// Create a map to store the population of different cities
	mapForDiffCities := map[string]int{
		"Oujda":   494252,
		"Kenitra": 431282,
		"Rabat":   577827,
	}
	// Print the map in raw form
	fmt.Println("Printing map in raw form:", mapForDiffCities)
	//Printing version for the map in the new map mapForDiffCitiesPrint
	fmt.Println("Printing the polished version of cities map: ")
	for cites, population := range mapForDiffCities {
		fmt.Printf("%s:%d\n", cites, population)
	}

	// Append through it to meet certain conditions
	// Create a new map to store big cities
	bigCities := make(map[string]int)
	for cities, population := range mapForDiffCities {
		if population > 450000 {
			bigCities[cities] = population
		}
	}

	newCities := map[string]int{

		"Casablanca": 2600000,
		"Marrakech":  900000,
	}

	//Merge the two maps
	for city, population := range mapForDiffCities {
		newCities[city] = population
	}
	fmt.Println(newCities)

	// Selecting only the big cities:

	metropolis := make(map[string]int)

	for city, population := range newCities {

		if population > 800000 {
			metropolis[city] = population
		}
	}
	//Print result
	fmt.Println("Citie(s) with a population bigger then 1 Million: ", metropolis)

	//Print result without the prefix map and []
	fmt.Println("Citie(s) with the largest population:")
	for printCity, printPopulation := range metropolis {
		fmt.Printf("%s with a population of:%d\n", printCity, printPopulation)
	}
}
