# 03 — Control Flow

`if`, `for`, `switch` — Go's three flow primitives. There is no `while`. There are no parentheses around conditions. There are no implicit fallthroughs.

## Run it

```bash
go run ./lessons/03-control-flow
```

## Test it

```bash
go test ./lessons/03-control-flow -v
```

## What to learn

- `for` is Go's only loop keyword. `for i := 0; i < n; i++`, `for cond`, and `for` (infinite) all use the same word.
- `switch` evaluates top-to-bottom. Cases **do not** fall through. Use `switch { case cond: ... }` (no expression) instead of an `if-else` ladder.
- Leap-year logic is a great early test of "ordering matters" in conditionals.

## Try yourself

1. Add `Days(year, month int) int` (handle leap February).
2. Refactor `FizzBuzz` to a single pass that builds the string and benchmark it against the current version with `go test -bench=.`.
