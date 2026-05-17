package main

import (
	"strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"", true},
		{"a", true},
		{"racecar", true},
		{"RaceCar", true},
		{"hello", false},
		{"A man, a plan, a canal: Panama", true},
		{"No lemon, no melon", true},
		{"12321", true},
		{"12345", false},
	}
	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			if got := IsPalindrome(tc.in); got != tc.want {
				t.Fatalf("IsPalindrome(%q) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	input := strings.Repeat("A man, a plan, a canal: Panama ", 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindrome(input)
	}
}

func BenchmarkIsPalindromeSimple(b *testing.B) {
	input := strings.Repeat("A man, a plan, a canal: Panama ", 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindromeSimple(input)
	}
}

// FuzzPalindrome lets the testing framework generate inputs and checks that
// reversing twice is identity-equivalent for the palindrome relation.
func FuzzPalindrome(f *testing.F) {
	for _, seed := range []string{"", "a", "racecar", "hello", "A B A"} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, s string) {
		// reversing s and re-checking should give the same answer
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		if IsPalindrome(s) != IsPalindrome(string(runes)) {
			t.Fatalf("palindrome status changed after reversing input: %q", s)
		}
	})
}
