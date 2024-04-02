package structures

import (
	"fmt"
	"time"
)

// Structs
// Adress Struct
type Adress struct {
	street, state, city, zipCode string
}

// - Define a struct representing a person with fields like name, age, and email. ✅
type Person struct {
	name, company, email, sex string
	age                       int
}

// - Define methods on the person struct to update its fields.
func (p *Person) gender(g string) {

	p.sex = g

}
func Structures() {
	actualYear := time.Now().Year()
	birthYear := 1995
	sex := "Male"

	//- Create instances of the person struct and print their details.✅
	person := Person{
		name:    "Atlas Kaisar",
		company: "KaisarInov8ors",
		email:   "atlas.kaisar@icloud.com",
		age:     actualYear - birthYear,
	}
	person.gender(sex)
	fmt.Printf("______Boss Data______\nName:%s\nCompany:%s\nEmail:%s\nAge:%d\nSex:%s\n", person.name, person.company, person.email, person.age, person.sex)

}
