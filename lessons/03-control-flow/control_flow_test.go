package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{5, "Buzz"},
		{9, "Fizz"},
		{10, "Buzz"},
		{15, "FizzBuzz"},
		{30, "FizzBuzz"},
		{-3, "Fizz"},
		{0, "FizzBuzz"}, // 0 is divisible by everything
	}
	for _, tc := range cases {
		if got := FizzBuzz(tc.in); got != tc.want {
			t.Errorf("FizzBuzz(%d) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

func TestIsLeapYear(t *testing.T) {
	cases := map[int]bool{
		1600: true,
		1700: false,
		1800: false,
		1900: false,
		2000: true,
		2024: true,
		2025: false,
	}
	for year, want := range cases {
		if got := IsLeapYear(year); got != want {
			t.Errorf("IsLeapYear(%d) = %v, want %v", year, got, want)
		}
	}
}

func TestGrade(t *testing.T) {
	cases := map[int]string{
		100: "A", 90: "A",
		89: "B", 80: "B",
		79: "C", 70: "C",
		69: "D", 60: "D",
		59: "F", 0: "F",
		-1: "?", 101: "?",
	}
	for percent, want := range cases {
		if got := Grade(percent); got != want {
			t.Errorf("Grade(%d) = %q, want %q", percent, got, want)
		}
	}
}

func TestFirstN(t *testing.T) {
	want := "1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz"
	if got := FirstN(15); got != want {
		t.Fatalf("FirstN(15) =\n  %q\nwant\n  %q", got, want)
	}
}
