# Challenges

Larger problems that combine multiple concepts from the [lessons](../lessons). Each one is a runnable package with its own `README.md` and tests.

| # | Challenge | Concepts |
|---|---|---|
| [01](01-fizzbuzz/) | FizzBuzz | control flow, `io.Writer`, table-driven tests |
| [02](02-linked-list/) | Doubly Linked List | pointers, generics, in-place mutation |
| [03](03-tail/) | `ztail` — simplified `tail -c N` | filesystem, `Seek`, argv parsing, exit codes |
| [04](04-library/) | Library Manager | structs, sentinel errors, `sync.RWMutex`, concurrency tests |

Run any one:

```bash
go run ./challenges/<dir>
go test ./challenges/<dir> -v
```

Or hit them all:

```bash
go test ./challenges/...
```
