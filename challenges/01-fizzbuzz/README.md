# Challenge 01 — FizzBuzz

The interview classic, written so it can actually be tested.

## Run

```bash
go run ./challenges/01-fizzbuzz
```

## Test

```bash
go test ./challenges/01-fizzbuzz -v
```

## The trick worth noticing

`FizzBuzz` takes an `io.Writer`, not stdout. Tests can pass a `bytes.Buffer` and assert on the exact output. This is the **dependency injection** pattern in its tiniest form, and it's how every testable Go program handles I/O.

## Extend it

- Make a generalized `Classify(n int, rules map[int]string) string`.
- Stream the output through a channel and have a goroutine print it.
