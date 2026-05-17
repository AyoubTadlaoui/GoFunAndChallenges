package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDivide(t *testing.T) {
	if q, err := Divide(6, 2); err != nil || q != 3 {
		t.Fatalf("Divide(6,2) = (%v, %v), want (3, nil)", q, err)
	}
	_, err := Divide(1, 0)
	if !errors.Is(err, ErrDivByZero) {
		t.Fatalf("Divide(1,0) err = %v, want ErrDivByZero", err)
	}
}

func TestReadFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "hello.txt")
	if err := os.WriteFile(path, []byte("hi"), 0o600); err != nil {
		t.Fatal(err)
	}
	got, err := ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile err = %v", err)
	}
	if got != "hi" {
		t.Fatalf("ReadFile = %q, want %q", got, "hi")
	}

	_, err = ReadFile(filepath.Join(dir, "missing"))
	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("missing file err = %v, want os.ErrNotExist", err)
	}
}

func TestLayerWrapping(t *testing.T) {
	err := Layer3()
	if err == nil {
		t.Fatal("Layer3 = nil, want wrapped error")
	}
	msg := err.Error()
	for _, want := range []string{"layer3", "layer2", "disk is on fire"} {
		if !strings.Contains(msg, want) {
			t.Errorf("missing %q in %q", want, msg)
		}
	}
}
