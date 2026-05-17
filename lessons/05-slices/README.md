# 05 — Slices

The workhorse collection type. Three words you need: **length**, **capacity**, **backing array**.

## Run it

```bash
go run ./lessons/05-slices
```

## Test it

```bash
go test ./lessons/05-slices -v
```

## What to learn

- A slice is a **header**: pointer, length, capacity. Multiple slices can share the same backing array — that's why `append` is sometimes a copy and sometimes not.
- `Filter` and `Map` here are tiny generic helpers (`[T any]`, `[T, U any]`). Generics arrived in Go 1.18 and they're how this kind of code stays type-safe.
- Allocate with `make([]T, 0, n)` when you know an upper bound — fewer reallocations.

## Try yourself

1. Add `Chunk[T any](s []T, size int) [][]T`.
2. Add `Reduce[T, U any](s []T, init U, f func(U, T) U) U`.
3. Make a deliberately broken `Filter` that returns `s[:0]` and reuses the input's array — observe how callers get surprised.
