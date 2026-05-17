package main

import (
	"reflect"
	"testing"
)

func TestSwap(t *testing.T) {
	a, b := 1, 2
	Swap(&a, &b)
	if a != 2 || b != 1 {
		t.Fatalf("Swap = (%d, %d), want (2, 1)", a, b)
	}
}

func TestInc(t *testing.T) {
	x := 41
	Inc(&x)
	if x != 42 {
		t.Fatalf("Inc = %d, want 42", x)
	}
}

func TestLinkedList_AppendAndValues(t *testing.T) {
	l := NewLinkedList()
	if l.Len() != 0 {
		t.Fatalf("empty Len = %d, want 0", l.Len())
	}
	l.Append(1)
	l.Append(2)
	l.Append(3)
	want := []int{1, 2, 3}
	if got := l.Values(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Values = %v, want %v", got, want)
	}
	if l.Len() != 3 {
		t.Fatalf("Len = %d, want 3", l.Len())
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	l := NewLinkedList()
	l.Append(2)
	l.Append(3)
	l.Prepend(1)
	want := []int{1, 2, 3}
	if got := l.Values(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Values = %v, want %v", got, want)
	}
}

func TestLinkedList_String(t *testing.T) {
	l := NewLinkedList()
	if l.String() != "(empty)" {
		t.Fatalf("empty String = %q, want (empty)", l.String())
	}
	l.Append(1)
	l.Append(2)
	l.Append(3)
	if l.String() != "1 -> 2 -> 3" {
		t.Fatalf("String = %q, want %q", l.String(), "1 -> 2 -> 3")
	}
}
