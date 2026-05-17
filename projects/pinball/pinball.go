// Package main is a small terminal pinball-style demo: a bouncing ball and a
// paddle you move with `w` and `s`. Press `q` to quit.
//
//	go run ./projects/pinball
//
// The game logic (Board, Step, MovePaddle) is plain data + functions so it
// can be unit-tested without a terminal.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	BoardWidth  = 30
	BoardHeight = 12
)

// Direction values for the ball's velocity components.
const (
	Up    = -1
	Down  = +1
	Left  = -1
	Right = +1
)

// Board holds the full mutable game state.
type Board struct {
	BallX, BallY int
	VelX, VelY   int
	PaddleY      int
}

// NewBoard returns the starting state.
func NewBoard() *Board {
	return &Board{
		BallX:   1,
		BallY:   BoardHeight / 2,
		VelX:    Right,
		VelY:    Down,
		PaddleY: BoardHeight / 2,
	}
}

// Step advances the ball one frame. The ball bounces off all four walls.
func (b *Board) Step() {
	b.BallX += b.VelX
	b.BallY += b.VelY
	if b.BallX <= 0 {
		b.BallX = 0
		b.VelX = Right
	} else if b.BallX >= BoardWidth-1 {
		b.BallX = BoardWidth - 1
		b.VelX = Left
	}
	if b.BallY <= 0 {
		b.BallY = 0
		b.VelY = Down
	} else if b.BallY >= BoardHeight-1 {
		b.BallY = BoardHeight - 1
		b.VelY = Up
	}
}

// MovePaddle moves the paddle by delta, clamped to the board.
func (b *Board) MovePaddle(delta int) {
	b.PaddleY += delta
	if b.PaddleY < 0 {
		b.PaddleY = 0
	}
	if b.PaddleY > BoardHeight-1 {
		b.PaddleY = BoardHeight - 1
	}
}

// Render returns the board as a string, one row per line.
func (b *Board) Render() string {
	var out strings.Builder
	paddleX := BoardWidth - 1
	for y := 0; y < BoardHeight; y++ {
		for x := 0; x < BoardWidth; x++ {
			switch {
			case x == b.BallX && y == b.BallY:
				out.WriteByte('*')
			case x == paddleX && y == b.PaddleY:
				out.WriteByte('|')
			default:
				out.WriteByte(' ')
			}
		}
		out.WriteByte('\n')
	}
	return out.String()
}

func main() {
	board := NewBoard()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Pinball — w/s to move paddle, q to quit, Enter to step.")
	for {
		fmt.Print("\033[H\033[2J") // clear screen (ANSI)
		fmt.Print(board.Render())
		fmt.Print("[w/s/q + Enter] > ")
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		switch strings.TrimSpace(line) {
		case "q":
			return
		case "w":
			board.MovePaddle(Up)
		case "s":
			board.MovePaddle(Down)
		}
		board.Step()
		time.Sleep(50 * time.Millisecond)
	}
}
