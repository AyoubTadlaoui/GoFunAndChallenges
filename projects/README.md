# Projects

Larger runnable Go programs that go beyond exercise scope.

| Project | What it is |
|---|---|
| [pinball](pinball/) | A tiny terminal ball-and-paddle demo. State is decoupled from I/O, so it's unit-testable. |

Run any one:

```bash
go run ./projects/pinball
```

## Add your own

Drop a new folder under `projects/`, put a `main.go` and a `README.md` in it, and you're done. The Makefile and CI auto-discover everything under `./...`.
