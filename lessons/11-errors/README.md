# 11 — Errors

`error` is just an interface with one method. Wrap with `%w`. Match with `errors.Is` / `errors.As`.

## Run / Test

```bash
go run ./lessons/11-errors
go test ./lessons/11-errors -v
```

## What to learn

- **Sentinel errors** (`var Err... = errors.New(...)`) let callers test with `errors.Is(err, ErrFoo)`.
- `fmt.Errorf("context: %w", err)` wraps so the chain is queryable via `errors.Is` / `errors.As`. Plain `%v` formats it but loses the chain — usually wrong.
- The previous version of this lesson had a real bug:

  ```go
  err := layer1              // captures the function value, not its result
  if err != nil { ... }      // function values are never nil — always true
  ```

  The fix is to **call** the function (`err := layer1()`) and check the returned error.

## Try yourself

1. Make a custom `type NotFoundError struct{ Path string }` that implements `Error() string` and `Is(target error) bool` to match `os.ErrNotExist`.
2. Add `errors.As` in the file reader so callers can inspect the underlying `*os.PathError`.
