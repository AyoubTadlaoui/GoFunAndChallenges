package pinball

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

const (
	boardWidth  = 20
	boardHeight = 10
)

type Ball struct {
	x, y     int
	velocity struct{ x, y int }
}

type Paddle struct {
	x, y int
}

func (b *Ball) move() {
	b.x += b.velocity.x
	b.y += b.velocity.y
	if b.x >= boardWidth || b.x <= 0 {
		b.velocity.x *= -1
	}
	if b.y >= boardHeight || b.y <= 0 {
		b.velocity.y *= -1
	}
}

func (p *Paddle) moveUp() {
	p.y--
	if p.y < 0 {
		p.y = 0
	}
}

func (p *Paddle) moveDown() {
	p.y++
	if p.y > boardHeight-1 {
		p.y = boardHeight - 1
	}
}

func getch() byte {
	var buf [1]byte
	syscall.Read(0, buf[:])
	return buf[0]
}

func Pinball() {
	board := make([][]byte, boardHeight)
	for i := range board {
		board[i] = make([]byte, boardWidth)
	}

	ball := Ball{x: 1, y: boardHeight / 2, velocity: struct{ x, y int }{1, 1}}
	paddle := Paddle{x: boardWidth - 1, y: boardHeight / 2}

	for {
		// Clear screen
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		// Move ball
		ball.move()

		// Draw board
		for y := range board {
			for x := range board[y] {
				if x == ball.x && y == ball.y {
					board[y][x] = '*'
				} else if x == paddle.x && y == paddle.y {
					board[y][x] = '|'
				} else {
					board[y][x] = ' '
				}
				fmt.Printf("%c", board[y][x])
			}
			fmt.Println()
		}

		// Move paddle based on user input
		var input byte
		fmt.Print("Press 'w' to move paddle up, 's' to move paddle down: ")
		input = getch()
		switch input {
		case 'w':
			paddle.moveUp()
		case 's':
			paddle.moveDown()
		}

		// Sleep for a short duration to control the speed
		time.Sleep(100 * time.Millisecond)
	}
}
