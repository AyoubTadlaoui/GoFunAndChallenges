// Package main is lesson 09: structs — Go's way to group fields.
//
//	go run ./lessons/09-structs
//	go test ./lessons/09-structs
package main

import "fmt"

// Address is a value type — small bundle of related fields.
type Address struct {
	City    string
	Region  string
	Country string
	Zip     int
}

// Person nests another struct (Address) and slices of strings (Languages).
type Person struct {
	Name      string
	Email     string
	Age       int
	Address   Address
	Languages []string
}

// String makes Person satisfy fmt.Stringer — `fmt.Println(p)` will use it.
func (p Person) String() string {
	return fmt.Sprintf("%s <%s>, age %d, %s, %s",
		p.Name, p.Email, p.Age, p.Address.City, p.Address.Country)
}

// HasLanguage reports whether the person lists lang in their languages.
func (p Person) HasLanguage(lang string) bool {
	for _, l := range p.Languages {
		if l == lang {
			return true
		}
	}
	return false
}

// WithEmail returns a copy of p with email replaced — illustrates the immutable
// "with" pattern. Useful when you want updates without surprising aliasing.
func (p Person) WithEmail(email string) Person {
	p.Email = email
	return p
}

// Adults returns the people in ps who are at least 18.
func Adults(ps []Person) []Person {
	out := make([]Person, 0, len(ps))
	for _, p := range ps {
		if p.Age >= 18 {
			out = append(out, p)
		}
	}
	return out
}
