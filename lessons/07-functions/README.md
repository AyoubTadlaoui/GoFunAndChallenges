# 07 — Functions

First-class. Multiple returns. Variadic. Closures.

## Run / Test

```bash
go run ./lessons/07-functions
go test ./lessons/07-functions -v
```

## What to learn

- **Multiple return values** are how Go does errors: `value, err`.
- A **variadic** parameter is just a slice in disguise: `nums ...int` inside is `[]int`.
- A **closure** captures variables from its enclosing scope **by reference**. Each call to `Counter` creates a fresh `n`.

## Try yourself

1. Add `Compose[A, B, C any](f func(B) C, g func(A) B) func(A) C`.
2. Make `DivMod` accept negative divisors and confirm Go's `%` returns a remainder with the sign of the dividend (different from Python!).
