package main

import "fmt"

func main() {
	ayoub := Person{
		Name:  "Atlas Kaisar",
		Email: "atlas@example.com",
		Age:   29,
		Address: Address{
			City:    "Oujda",
			Region:  "Oriental",
			Country: "Morocco",
			Zip:     60000,
		},
		Languages: []string{"Arabic", "French", "English", "Go"},
	}

	fmt.Println("=== Structs ===")
	fmt.Println(ayoub)
	fmt.Println("Speaks Go?  ", ayoub.HasLanguage("Go"))
	fmt.Println("Speaks Rust?", ayoub.HasLanguage("Rust"))

	updated := ayoub.WithEmail("ayoub@example.com")
	fmt.Println("Updated:", updated)
	fmt.Println("Original unchanged:", ayoub.Email)

	people := []Person{ayoub, {Name: "Kid", Age: 10}, {Name: "Boss", Age: 45}}
	fmt.Printf("Adults: %d of %d\n", len(Adults(people)), len(people))
}
