# 14 — Testing, Benchmarks, Fuzzing

Go ships with a first-class testing framework. No external library needed.

## Try every mode

```bash
# Unit tests
go test ./lessons/14-testing -v

# Benchmarks (compare two implementations)
go test ./lessons/14-testing -bench=. -benchmem

# Fuzz for 5 seconds
go test ./lessons/14-testing -fuzz=FuzzPalindrome -fuzztime=5s
```

## What to learn

- **Table-driven tests** — define a slice of cases, loop with `t.Run` so each subtest shows up by name.
- **Benchmarks** live in `BenchmarkXxx(b *testing.B)` and loop `b.N` times. Use `b.ResetTimer()` after setup.
- **Fuzzing** (`FuzzXxx(f *testing.F)`) auto-generates inputs from seeds and the corpus. Great at finding edge cases you didn't think of.

## Try yourself

1. Add `BenchmarkParallelSquare` against the concurrency lesson at several worker counts.
2. Wire `go test -coverprofile=coverage.out ./...` into your workflow, then `go tool cover -html=coverage.out`.
