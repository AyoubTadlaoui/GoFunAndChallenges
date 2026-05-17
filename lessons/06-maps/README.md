# 06 — Maps

Hash tables, Go-style.

## Run / Test

```bash
go run ./lessons/06-maps
go test ./lessons/06-maps -v
```

## What to learn

- `make(map[K]V, hint)` reserves capacity — useful when you know the size.
- Lookup with comma-ok: `v, ok := m[k]`. `ok == false` means the key is absent. `v` is the zero value either way.
- Iteration order is **randomized** intentionally. Don't rely on it. Sort the keys when you need determinism (see `TopN`).

## Try yourself

1. Replace the body of `TopN` with the new `slices.SortFunc` from the standard library.
2. Write `Invert(m map[string]int) map[int][]string` — values become keys, keys collected per value.
