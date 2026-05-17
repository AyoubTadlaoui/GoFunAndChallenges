package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestFizzBuzz_FirstFifteen(t *testing.T) {
	var buf bytes.Buffer
	if err := FizzBuzz(&buf, 15); err != nil {
		t.Fatal(err)
	}
	want := "1\n2\nFizz\n4\nBuzz\nFizz\n7\n8\nFizz\nBuzz\n11\nFizz\n13\n14\nFizzBuzz\n"
	if buf.String() != want {
		t.Fatalf("FizzBuzz(15) =\n%q\nwant\n%q", buf.String(), want)
	}
}

func TestFizzBuzz_Zero(t *testing.T) {
	var buf bytes.Buffer
	if err := FizzBuzz(&buf, 0); err != nil {
		t.Fatal(err)
	}
	if buf.Len() != 0 {
		t.Fatalf("FizzBuzz(0) produced output: %q", buf.String())
	}
}

func TestFizzBuzz_LineCount(t *testing.T) {
	var buf bytes.Buffer
	if err := FizzBuzz(&buf, 100); err != nil {
		t.Fatal(err)
	}
	lines := strings.Count(buf.String(), "\n")
	if lines != 100 {
		t.Fatalf("FizzBuzz(100) wrote %d lines, want 100", lines)
	}
}
