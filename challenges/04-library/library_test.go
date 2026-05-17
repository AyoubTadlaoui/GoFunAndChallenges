package main

import (
	"errors"
	"reflect"
	"sync"
	"testing"
)

func seedLibrary(t *testing.T) *Library {
	t.Helper()
	l := New()
	mustAdd := func(b Book) {
		if err := l.Add(b); err != nil {
			t.Fatal(err)
		}
	}
	mustAdd(Book{ID: "1", Title: "The Go Programming Language", Author: "Donovan", Year: 2015})
	mustAdd(Book{ID: "2", Title: "Learning Go", Author: "Bodner", Year: 2021})
	mustAdd(Book{ID: "3", Title: "Learning Go 2e", Author: "Bodner", Year: 2024})
	return l
}

func TestAdd_Duplicate(t *testing.T) {
	l := New()
	if err := l.Add(Book{ID: "1"}); err != nil {
		t.Fatal(err)
	}
	err := l.Add(Book{ID: "1"})
	if !errors.Is(err, ErrDuplicateBook) {
		t.Fatalf("err = %v, want ErrDuplicateBook", err)
	}
}

func TestLendAndReturn(t *testing.T) {
	l := seedLibrary(t)

	if err := l.Lend("1", "Atlas"); err != nil {
		t.Fatal(err)
	}
	b, err := l.Get("1")
	if err != nil || b.Borrower != "Atlas" {
		t.Fatalf("after Lend, book = %+v, err = %v", b, err)
	}

	err = l.Lend("1", "Mehdi")
	if !errors.Is(err, ErrAlreadyOnLoan) {
		t.Fatalf("second Lend err = %v, want ErrAlreadyOnLoan", err)
	}

	if err := l.Return("1"); err != nil {
		t.Fatal(err)
	}
	if err := l.Return("1"); !errors.Is(err, ErrNotOnLoan) {
		t.Fatalf("second Return err = %v, want ErrNotOnLoan", err)
	}
}

func TestErrors_NotFound(t *testing.T) {
	l := New()
	if _, err := l.Get("nope"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("Get err = %v, want ErrNotFound", err)
	}
	if err := l.Lend("nope", "u"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("Lend err = %v, want ErrNotFound", err)
	}
	if err := l.Return("nope"); !errors.Is(err, ErrNotFound) {
		t.Fatalf("Return err = %v, want ErrNotFound", err)
	}
}

func TestByAuthor_SortedByYear(t *testing.T) {
	l := seedLibrary(t)
	got := l.ByAuthor("Bodner")
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
	if got[0].Year != 2021 || got[1].Year != 2024 {
		t.Fatalf("years = %d, %d, want 2021, 2024", got[0].Year, got[1].Year)
	}
}

func TestAvailableTitlesExcludesBorrowed(t *testing.T) {
	l := seedLibrary(t)
	if err := l.Lend("2", "Atlas"); err != nil {
		t.Fatal(err)
	}
	want := []string{"Learning Go 2e", "The Go Programming Language"}
	if got := l.AvailableTitles(); !reflect.DeepEqual(got, want) {
		t.Fatalf("AvailableTitles = %v, want %v", got, want)
	}
}

// TestConcurrentAccess hammers the library with concurrent Lend / Return and
// then verifies the final state is consistent.
func TestConcurrentAccess(t *testing.T) {
	l := New()
	const N = 100
	for i := 0; i < N; i++ {
		if err := l.Add(Book{ID: idOf(i), Title: "T"}); err != nil {
			t.Fatal(err)
		}
	}
	var wg sync.WaitGroup
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_ = l.Lend(idOf(i), "u")
			_ = l.Return(idOf(i))
		}(i)
	}
	wg.Wait()
	if got := l.AvailableTitles(); len(got) != N {
		t.Fatalf("after concurrent traffic, available = %d, want %d", len(got), N)
	}
}

func idOf(i int) string {
	const digits = "0123456789"
	if i == 0 {
		return "0"
	}
	var out []byte
	for i > 0 {
		out = append([]byte{digits[i%10]}, out...)
		i /= 10
	}
	return string(out)
}
