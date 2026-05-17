// Package main is challenge 02: doubly linked list with full operations.
//
//	go run ./challenges/02-linked-list
//	go test ./challenges/02-linked-list
//
// Compared to the singly linked list in lesson 10, this version supports O(1)
// removal at any node, bidirectional traversal, and constant-time Len.
package main

import (
	"fmt"
	"strings"
)

// Node is one element of a doubly linked list.
type Node[T any] struct {
	Value      T
	prev, next *Node[T]
}

// List is a doubly linked list parameterized by T.
type List[T any] struct {
	head, tail *Node[T]
	length     int
}

// New returns an empty list.
func New[T any]() *List[T] { return &List[T]{} }

// Len returns the number of elements.
func (l *List[T]) Len() int { return l.length }

// PushBack appends v at the tail. Returns the new node.
func (l *List[T]) PushBack(v T) *Node[T] {
	n := &Node[T]{Value: v, prev: l.tail}
	if l.tail != nil {
		l.tail.next = n
	} else {
		l.head = n
	}
	l.tail = n
	l.length++
	return n
}

// PushFront prepends v at the head. Returns the new node.
func (l *List[T]) PushFront(v T) *Node[T] {
	n := &Node[T]{Value: v, next: l.head}
	if l.head != nil {
		l.head.prev = n
	} else {
		l.tail = n
	}
	l.head = n
	l.length++
	return n
}

// Remove detaches n from l. n must belong to l (the caller's responsibility —
// we don't verify, just like container/list).
func (l *List[T]) Remove(n *Node[T]) T {
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		l.head = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	} else {
		l.tail = n.prev
	}
	n.prev, n.next = nil, nil
	l.length--
	return n.Value
}

// Values returns all elements head-to-tail.
func (l *List[T]) Values() []T {
	out := make([]T, 0, l.length)
	for n := l.head; n != nil; n = n.next {
		out = append(out, n.Value)
	}
	return out
}

// Reverse reverses the list in place. O(n), no allocations.
func (l *List[T]) Reverse() {
	for n := l.head; n != nil; {
		n.prev, n.next = n.next, n.prev
		n = n.prev // we just swapped, so prev is the old next
	}
	l.head, l.tail = l.tail, l.head
}

// String renders the list using fmt formatting on each element.
func (l *List[T]) String() string {
	parts := make([]string, 0, l.length)
	for n := l.head; n != nil; n = n.next {
		parts = append(parts, fmt.Sprintf("%v", n.Value))
	}
	if len(parts) == 0 {
		return "(empty)"
	}
	return strings.Join(parts, " <-> ")
}
