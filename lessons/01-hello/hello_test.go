package main

import "testing"

func TestGreet(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"named", "Ayoub", "Hello, Ayoub! Welcome to Go."},
		{"empty falls back to world", "", "Hello, world! Welcome to Go."},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Greet(tc.in); got != tc.want {
				t.Fatalf("Greet(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
