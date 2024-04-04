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
	actualYear := time.Now().Year()
	birthYear := 1995
	sex := "Male"
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
		name:    "Atlas Kaisar",
		company: "KaisarInov8ors",
		email:   "atlas.kaisar@icloud.com",
		age:     actualYear - birthYear,
		adress:  adress,
		langs:   mainLangs,
	}

	person.gender(sex)
	fmt.Printf("______Boss Data______\nName:%s\nCompany:%s\nEmail:%s\nAge:%d\nSex:%s\n%+v\n%+v\n", person.name, person.company, person.email, person.age, person.sex, person.adress, person.langs)

	fmt.Println("===Adress===")
	fmt.Println("", adress.city, "\n", adress.region, "\n", adress.country, "\n", adress.zipcode)
	fmt.Printf("%+v\n", adress)
}
