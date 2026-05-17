# 13 — Concurrency

Goroutines. Channels. Mutex. Atomic.

## Run / Test

```bash
go run ./lessons/13-concurrency
go test ./lessons/13-concurrency -race -v
```

The `-race` flag is the **single most useful flag in Go**. It instruments your binary to detect data races. Always run it on concurrent code.

## What to learn

- A **goroutine** is a function call prefixed with `go`. Cheap (kilobytes of stack), but they still need to be coordinated — usually with channels or `sync.WaitGroup`.
- A **channel** typed `chan T` is a typed pipe. `close(ch)` signals "no more values".
- A **worker pool** (here `ParallelSquare`) sends jobs into a channel and spawns N consumers — bounded parallelism without spawning a goroutine per task.
- `sync.Mutex` for arbitrary critical sections. `sync/atomic` for single-word counters.

## Try yourself

1. Replace the mutex in `SafeCounter` with `atomic.Int64` and benchmark both with `go test -bench=. -benchmem`.
2. Add a `ctx context.Context` parameter to `ParallelSquare` and stop early when the context is canceled.
3. Build a fan-in: many producers writing to one channel, one consumer draining it.
