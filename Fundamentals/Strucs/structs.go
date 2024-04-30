package structures

import (
	"fmt"
	"time"
)

// Structs

// - Define a struct representing a person with fields like name, age, and email. ✅
type Person struct {
	name, company, email, sex string
	age                       int
	adress                    Adress
	langs                     MainLanguages
}

// Adress Struct
type Adress struct {
	city, region, country string
	zipcode               int
}

// Natural Languages languages
type MainLanguages struct {
	naturalLanguages, programminglanguages []string
}

type AllLanguages struct {
	naturalLanguages, programminglanguages MainLanguages
}

//

// -Method to initiate the values of Adress Struct
func (a *Adress) adressMethod(city, region, country string, zip int) {
	a.city = city
	a.region = region
	a.country = country
	a.zipcode = zip
}

// -Method to initiate values of all languages:
func (mainLan *MainLanguages) languageStack(nattyLang, progLang []string) {
	mainLan.naturalLanguages = nattyLang
	mainLan.programminglanguages = progLang
}

// - Define methods on the person struct to update its fields.
func (p *Person) gender(g string) {

	p.sex = g

}
func Structures() {
	var name, email, gender string
	fmt.Println("Enter your name")
	fmt.Scan(&name)
	fmt.Println("Enter Your Email")
	fmt.Scan(&email)
	actualYear := time.Now().Year()
	var birthYear int
	fmt.Println("Enter Year you were Born")
	fmt.Scan(&birthYear)
	fmt.Println("Are you Male or Female?\nPress m for male.\nPress f for female.")
	fmt.Scan(&gender)
	if gender == "m" {
		gender = "Male"
	} else if gender == "f" {
		gender = "Female"
	} else {
		gender = "Undefined"
	}
	//Languages:
	nattylangs := [...]string{"English", "French", "Arabic"}
	proglangs := []string{"Go", "Dart"}
	mainLangs := MainLanguages{}
	mainLangs.languageStack(nattylangs[:], proglangs)

	//Adress
	city, region, country := "Oujda", "Oriental Metropolitan Aria", "Morocco"
	zip := 60000
	adress := Adress{}
	adress.adressMethod(city, region, country, zip)
	//- Create instances of the person struct and print their details.✅
	person := Person{
		name:    name,
		company: "Zone01",
		email:   email,
		age:     actualYear - birthYear,
		adress:  adress,
		langs:   mainLangs,
	}

	person.gender(gender)
	fmt.Printf("______Boss Data______\nName:%s\nCompany:%s\nEmail:%s\nAge:%d\nSex:%s\n%+v\n%+v\n", person.name, person.company, person.email, person.age, person.sex, person.adress, person.langs)

	fmt.Println("===Adress===")
	fmt.Println("", adress.city, "\n", adress.region, "\n", adress.country, "\n", adress.zipcode)
	fmt.Printf("%+v\n", adress)
}
