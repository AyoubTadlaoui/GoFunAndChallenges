// Package main is lesson 10: pointers and a tiny linked list.
//
//	go run ./lessons/10-pointers
//	go test ./lessons/10-pointers
//
// Why pointers? To mutate a caller's value, share a big struct without copying,
// or build data structures that link to themselves (lists, trees).
package main

import (
	"fmt"
	"strings"
)

// Swap exchanges the values that a and b point to.
func Swap(a, b *int) {
	*a, *b = *b, *a
}

// Inc increments the int pointed to by p.
func Inc(p *int) {
	*p++
}

// Node holds one int and a link to the next node (or nil at the tail).
type Node struct {
	Value int
	Next  *Node
}

// LinkedList is a singly-linked list that supports Append (tail-insert) and
// Prepend (head-insert). It tracks tail so Append stays O(1).
type LinkedList struct {
	head, tail *Node
	length     int
}

// NewLinkedList returns an empty list.
func NewLinkedList() *LinkedList { return &LinkedList{} }

// Len returns the number of nodes.
func (l *LinkedList) Len() int { return l.length }

// Append adds v at the tail in O(1).
func (l *LinkedList) Append(v int) {
	n := &Node{Value: v}
	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		l.tail.Next = n
		l.tail = n
	}
	l.length++
}

// Prepend adds v at the head in O(1).
func (l *LinkedList) Prepend(v int) {
	n := &Node{Value: v, Next: l.head}
	l.head = n
	if l.tail == nil {
		l.tail = n
	}
	l.length++
}

// Values returns the values from head to tail.
func (l *LinkedList) Values() []int {
	out := make([]int, 0, l.length)
	for n := l.head; n != nil; n = n.Next {
		out = append(out, n.Value)
	}
	return out
}

// String renders the list as "a -> b -> c".
func (l *LinkedList) String() string {
	parts := make([]string, 0, l.length)
	for n := l.head; n != nil; n = n.Next {
		parts = append(parts, fmt.Sprintf("%d", n.Value))
	}
	if len(parts) == 0 {
		return "(empty)"
	}
	return strings.Join(parts, " -> ")
}
