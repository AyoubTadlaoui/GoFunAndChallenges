// Package main is challenge 04: a tiny library / borrowing system.
//
//	go run ./challenges/04-library
//	go test ./challenges/04-library
//
// The challenge in the original repo: "Design a program to manage a library
// with books and their availability. Implement functions to add new books,
// lend books to borrowers, and track book availability."
package main

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

// Errors returned by the library API.
var (
	ErrNotFound      = errors.New("library: book not found")
	ErrAlreadyOnLoan = errors.New("library: book already on loan")
	ErrNotOnLoan     = errors.New("library: book is not currently on loan")
	ErrDuplicateBook = errors.New("library: book with this ID already exists")
)

// Book describes a single title in the catalog.
type Book struct {
	ID       string
	Title    string
	Author   string
	Year     int
	Borrower string // empty when available
}

// Available reports whether the book is currently on a shelf.
func (b Book) Available() bool { return b.Borrower == "" }

// Library is a concurrency-safe catalog. Multiple goroutines may add, lend,
// and return books without external synchronization.
type Library struct {
	mu    sync.RWMutex
	books map[string]*Book
}

// New returns an empty library.
func New() *Library {
	return &Library{books: map[string]*Book{}}
}

// Add registers a book. Returns ErrDuplicateBook if an ID is reused.
func (l *Library) Add(b Book) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if _, ok := l.books[b.ID]; ok {
		return fmt.Errorf("%w: id=%q", ErrDuplicateBook, b.ID)
	}
	bb := b // copy to avoid sharing with the caller
	l.books[b.ID] = &bb
	return nil
}

// Lend marks a book as borrowed by user.
func (l *Library) Lend(id, user string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	b, ok := l.books[id]
	if !ok {
		return fmt.Errorf("%w: id=%q", ErrNotFound, id)
	}
	if !b.Available() {
		return fmt.Errorf("%w: id=%q borrower=%q", ErrAlreadyOnLoan, id, b.Borrower)
	}
	b.Borrower = user
	return nil
}

// Return marks a book as returned and back on the shelf.
func (l *Library) Return(id string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	b, ok := l.books[id]
	if !ok {
		return fmt.Errorf("%w: id=%q", ErrNotFound, id)
	}
	if b.Available() {
		return fmt.Errorf("%w: id=%q", ErrNotOnLoan, id)
	}
	b.Borrower = ""
	return nil
}

// Get returns a copy of the book with the given ID.
func (l *Library) Get(id string) (Book, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()
	b, ok := l.books[id]
	if !ok {
		return Book{}, fmt.Errorf("%w: id=%q", ErrNotFound, id)
	}
	return *b, nil
}

// ByAuthor returns every book by author, sorted by year ascending.
func (l *Library) ByAuthor(author string) []Book {
	l.mu.RLock()
	defer l.mu.RUnlock()
	out := []Book{}
	for _, b := range l.books {
		if b.Author == author {
			out = append(out, *b)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Year < out[j].Year })
	return out
}

// AvailableTitles returns the titles of all available books, sorted.
func (l *Library) AvailableTitles() []string {
	l.mu.RLock()
	defer l.mu.RUnlock()
	titles := []string{}
	for _, b := range l.books {
		if b.Available() {
			titles = append(titles, b.Title)
		}
	}
	sort.Strings(titles)
	return titles
}
