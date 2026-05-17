package main

import (
	"strings"
	"testing"
)

func TestStep_BouncesOffWalls(t *testing.T) {
	b := NewBoard()
	// Force the ball into the bottom-right corner heading further into the wall.
	b.BallX, b.BallY = BoardWidth-1, BoardHeight-1
	b.VelX, b.VelY = Right, Down
	b.Step()
	if b.VelX != Left {
		t.Fatalf("VelX after right wall = %d, want Left(%d)", b.VelX, Left)
	}
	if b.VelY != Up {
		t.Fatalf("VelY after bottom wall = %d, want Up(%d)", b.VelY, Up)
	}
}

func TestMovePaddle_Clamps(t *testing.T) {
	b := NewBoard()
	for i := 0; i < 100; i++ {
		b.MovePaddle(Up)
	}
	if b.PaddleY != 0 {
		t.Fatalf("PaddleY = %d, want 0", b.PaddleY)
	}
	for i := 0; i < 100; i++ {
		b.MovePaddle(Down)
	}
	if b.PaddleY != BoardHeight-1 {
		t.Fatalf("PaddleY = %d, want %d", b.PaddleY, BoardHeight-1)
	}
}

func TestRender_ContainsBallAndPaddle(t *testing.T) {
	b := NewBoard()
	out := b.Render()
	if !strings.Contains(out, "*") {
		t.Fatal("rendered board missing ball '*'")
	}
	if !strings.Contains(out, "|") {
		t.Fatal("rendered board missing paddle '|'")
	}
	lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
	if len(lines) != BoardHeight {
		t.Fatalf("rendered %d lines, want %d", len(lines), BoardHeight)
	}
	for i, line := range lines {
		if len(line) != BoardWidth {
			t.Fatalf("line %d width = %d, want %d", i, len(line), BoardWidth)
		}
	}
}
