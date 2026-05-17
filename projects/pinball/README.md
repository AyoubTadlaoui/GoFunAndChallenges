# Pinball

A tiny terminal pinball-style demo: a ball bouncing in a box, with a paddle you can push around.

## Play

```bash
go run ./projects/pinball
```

Controls: `w` then Enter (paddle up), `s` then Enter (paddle down), `q` then Enter (quit).

> Why "press Enter"? Reading single keystrokes portably requires putting the terminal into raw mode, which is fiddly across macOS/Linux/Windows. The line-buffered version keeps the lesson clear and the code testable.

## Test

```bash
go test ./projects/pinball -v
```

## What it shows

- **Decouple state from I/O**: `Board.Step`, `Board.MovePaddle`, and `Board.Render` are pure data + functions. The `main` loop is the only thing that touches the terminal — that's why everything except `main` is unit-testable.
- ANSI escape `\033[H\033[2J` clears the screen on every frame.

## Extend it

- Detect "you missed!" when the ball reaches the right edge on a row where the paddle isn't.
- Speed up the ball every 10 successful hits.
- Swap stdin for `golang.org/x/term` raw mode and play in real time.
