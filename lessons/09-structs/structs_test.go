package main

import "testing"

func sample() Person {
	return Person{
		Name:      "Ayoub",
		Email:     "a@b.com",
		Age:       29,
		Address:   Address{City: "Oujda", Country: "Morocco", Zip: 60000},
		Languages: []string{"Arabic", "Go"},
	}
}

func TestHasLanguage(t *testing.T) {
	p := sample()
	if !p.HasLanguage("Go") {
		t.Fatal("HasLanguage(Go) = false, want true")
	}
	if p.HasLanguage("Rust") {
		t.Fatal("HasLanguage(Rust) = true, want false")
	}
}

func TestWithEmail(t *testing.T) {
	p := sample()
	q := p.WithEmail("new@example.com")
	if q.Email != "new@example.com" {
		t.Fatalf("new email = %q, want new@example.com", q.Email)
	}
	if p.Email == q.Email {
		t.Fatal("WithEmail mutated the original")
	}
}

func TestStringer(t *testing.T) {
	got := sample().String()
	if got == "" {
		t.Fatal("Stringer returned empty")
	}
}

func TestAdults(t *testing.T) {
	got := Adults([]Person{
		{Name: "A", Age: 10},
		{Name: "B", Age: 18},
		{Name: "C", Age: 99},
	})
	if len(got) != 2 || got[0].Name != "B" || got[1].Name != "C" {
		t.Fatalf("Adults = %v, want B and C", got)
	}
}
