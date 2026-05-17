package main

import (
	"reflect"
	"testing"
)

func TestPushBack(t *testing.T) {
	l := New[int]()
	for i := 1; i <= 3; i++ {
		l.PushBack(i)
	}
	if got, want := l.Values(), []int{1, 2, 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Values = %v, want %v", got, want)
	}
	if l.Len() != 3 {
		t.Fatalf("Len = %d, want 3", l.Len())
	}
}

func TestPushFront(t *testing.T) {
	l := New[string]()
	l.PushBack("b")
	l.PushBack("c")
	l.PushFront("a")
	if got, want := l.Values(), []string{"a", "b", "c"}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Values = %v, want %v", got, want)
	}
}

func TestRemoveMiddle(t *testing.T) {
	l := New[int]()
	l.PushBack(1)
	mid := l.PushBack(2)
	l.PushBack(3)
	if v := l.Remove(mid); v != 2 {
		t.Fatalf("Remove returned %d, want 2", v)
	}
	if got, want := l.Values(), []int{1, 3}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Values = %v, want %v", got, want)
	}
	if l.Len() != 2 {
		t.Fatalf("Len = %d, want 2", l.Len())
	}
}

func TestRemoveHeadAndTail(t *testing.T) {
	l := New[int]()
	a := l.PushBack(1)
	b := l.PushBack(2)
	c := l.PushBack(3)
	l.Remove(a) // head
	l.Remove(c) // tail
	if got, want := l.Values(), []int{2}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Values = %v, want %v", got, want)
	}
	l.Remove(b)
	if l.Len() != 0 {
		t.Fatalf("Len after all removed = %d, want 0", l.Len())
	}
}

func TestReverse(t *testing.T) {
	l := New[int]()
	for _, v := range []int{1, 2, 3, 4, 5} {
		l.PushBack(v)
	}
	l.Reverse()
	if got, want := l.Values(), []int{5, 4, 3, 2, 1}; !reflect.DeepEqual(got, want) {
		t.Fatalf("Reverse = %v, want %v", got, want)
	}
	// reverse twice -> back to original
	l.Reverse()
	if got, want := l.Values(), []int{1, 2, 3, 4, 5}; !reflect.DeepEqual(got, want) {
		t.Fatalf("double-reverse = %v, want %v", got, want)
	}
}

func TestEmpty(t *testing.T) {
	l := New[int]()
	if l.String() != "(empty)" {
		t.Fatalf("empty String = %q, want (empty)", l.String())
	}
	l.Reverse() // must not panic
	if l.Len() != 0 {
		t.Fatalf("Len = %d, want 0", l.Len())
	}
}
