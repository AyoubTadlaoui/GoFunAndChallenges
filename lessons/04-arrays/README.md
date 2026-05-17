# 04 — Arrays

Fixed-size, value-typed sequences.

## Run it

```bash
go run ./lessons/04-arrays
```

## Test it

```bash
go test ./lessons/04-arrays -v
```

## What to learn

- `[5]int` and `[6]int` are **different types**. The length is part of the type.
- Arrays are **value types** — passing one to a function copies it. Mutations inside the function don't reach the caller. (See the printed proof in `main`.)
- For dynamic sizes, almost everything you actually want is a slice — covered next.

## Try yourself

1. Make `MinMax` generic over any ordered type (`[T cmp.Ordered, N int](a [N]T)`).
2. Add `Contains(a [5]int, v int) bool`.
