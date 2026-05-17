# GoFunAndChallenges

> A hands-on Go course. **14 runnable lessons**, **4 multi-concept challenges**, and **1 project** — each with code, README, and tests you can run today.

[![CI](https://github.com/AyoubTadlaoui/GoFunAndChallenges/actions/workflows/ci.yml/badge.svg)](https://github.com/AyoubTadlaoui/GoFunAndChallenges/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/AyoubTadlaoui/GoFunAndChallenges.svg)](https://pkg.go.dev/github.com/AyoubTadlaoui/GoFunAndChallenges)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

---

## Quickstart (under 60 seconds)

```bash
# 1. Install Go ≥ 1.22 — https://go.dev/dl/
go version

# 2. Clone
git clone https://github.com/AyoubTadlaoui/GoFunAndChallenges.git
cd GoFunAndChallenges

# 3. Run lesson 01
go run ./lessons/01-hello

# 4. Run every test in the course
go test ./...
```

That's it. Every lesson is a runnable package — no commented-out imports, no "uncomment this to enable that". Pick a lesson, run it, test it, edit it.

---

## What's inside

```
GoFunAndChallenges/
├── lessons/        14 progressive lessons (hello → testing)
├── challenges/      4 multi-concept exercises
├── projects/        Larger programs (currently: pinball)
├── docs/            CURRICULUM.md — the learning path
├── .github/         CI workflow
├── Makefile         make help for the full list of tasks
├── .golangci.yml    Lint config
└── go.mod
```

### Lessons

A guided tour of the language. Each one has `main.go`, `*_test.go`, and a per-lesson `README.md`.

| # | Topic | Run |
|---|---|---|
| [01](lessons/01-hello)            | Hello, Go                    | `go run ./lessons/01-hello` |
| [02](lessons/02-variables)        | Variables & constants        | `go run ./lessons/02-variables` |
| [03](lessons/03-control-flow)     | `if` / `for` / `switch`      | `go run ./lessons/03-control-flow` |
| [04](lessons/04-arrays)           | Fixed-size arrays            | `go run ./lessons/04-arrays` |
| [05](lessons/05-slices)           | Slices (with generics)       | `go run ./lessons/05-slices` |
| [06](lessons/06-maps)             | Maps                         | `go run ./lessons/06-maps` |
| [07](lessons/07-functions)        | Functions, variadic, closures| `go run ./lessons/07-functions` |
| [08](lessons/08-methods)          | Methods & interfaces         | `go run ./lessons/08-methods` |
| [09](lessons/09-structs)          | Structs & `fmt.Stringer`     | `go run ./lessons/09-structs` |
| [10](lessons/10-pointers)         | Pointers & a linked list     | `go run ./lessons/10-pointers` |
| [11](lessons/11-errors)           | Errors, `%w`, `errors.Is`    | `go run ./lessons/11-errors` |
| [12](lessons/12-packages)         | Packages & modules           | `go run ./lessons/12-packages` |
| [13](lessons/13-concurrency)      | Goroutines, channels, mutex  | `go run ./lessons/13-concurrency` |
| [14](lessons/14-testing)          | Tests, benchmarks, fuzzing   | `go run ./lessons/14-testing` |

### Challenges

Bigger problems that combine multiple concepts.

| # | Challenge | Highlights |
|---|---|---|
| [01](challenges/01-fizzbuzz)     | FizzBuzz with `io.Writer`        | Dependency injection for testability |
| [02](challenges/02-linked-list)  | Generic doubly linked list       | `List[T]`, O(1) remove, in-place reverse |
| [03](challenges/03-tail)         | `ztail` — simplified `tail -c N` | `Seek`, multi-file, exit codes |
| [04](challenges/04-library)      | Concurrent library catalog       | Sentinel errors, `sync.RWMutex`, `-race` test |

### Projects

| Project | What it is |
|---|---|
| [pinball](projects/pinball) | Terminal ball-and-paddle. State and I/O are split, so the game logic is unit-tested. |

See [`docs/CURRICULUM.md`](docs/CURRICULUM.md) for the suggested study order and time budget.

---

## Common tasks

```bash
make help        # list everything
make check       # fmt + vet + test  (the default sanity check)
make test        # go test ./...
make test-race   # go test -race ./...    (always use this on concurrent code)
make cover       # coverage profile + HTML report
make bench       # run all benchmarks
make lint        # golangci-lint run
```

---

## Why this layout

The earlier version of the repo asked learners to uncomment imports in a central `main.go` to switch between exercises. That was a friction trap: imports broke, package names drifted, and tests didn't exist.

This rewrite trades that workflow for the standard Go convention:

- **One directory = one runnable / testable package.** `go run ./lessons/05-slices` just works.
- **Every lesson has tests.** That's what `go test ./...` is for — and CI proves it on every push.
- **Pure logic is testable, I/O lives at the edges.** See `challenges/01-fizzbuzz` and `projects/pinball` for the pattern in miniature.
- **Bugs from the original have been fixed.** `math.Phi` → `math.Pi`, broken error wrapping, "surface" misnamed as circumference, etc.

---

## Contributing

PRs welcome. See [CONTRIBUTING.md](CONTRIBUTING.md) for the workflow.

## License

MIT — see [LICENSE](LICENSE).

---

### About the author

I'm **Ayoub Tadlaoui** — *Atlas Kaisar* — a problem-solver from Morocco, building software since 2016.
My journey into programming started with **modulo topics and abelian groups in mathematics**, where I discovered the power of computational thinking.

> "High performance knows no part-time commitment."
