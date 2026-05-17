package calc

import (
	"errors"
	"testing"
)

func TestArithmetic(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Errorf("Add = %v, want 5", got)
	}
	if got := Sub(5, 2); got != 3 {
		t.Errorf("Sub = %v, want 3", got)
	}
	if got := Mul(4, 5); got != 20 {
		t.Errorf("Mul = %v, want 20", got)
	}
}

func TestDiv(t *testing.T) {
	q, err := Div(10, 4)
	if err != nil || q != 2.5 {
		t.Fatalf("Div(10,4) = (%v, %v), want (2.5, nil)", q, err)
	}
	if _, err := Div(1, 0); !errors.Is(err, ErrDivByZero) {
		t.Fatalf("Div(1,0) err = %v, want ErrDivByZero", err)
	}
}
