# Curriculum

A suggested path through the course. Total time: **8 – 12 focused hours** if you do the "Try yourself" sections at the end of each lesson.

## Week 1 — Language basics (~3 hours)

| Lesson | Focus | Key idea |
|---|---|---|
| [01 — Hello](../lessons/01-hello)            | Run your first program | Keep logic out of `main` so you can test it |
| [02 — Variables](../lessons/02-variables)    | `var`, `:=`, `const`, type switch | `math.Pi` ≠ `math.Phi` |
| [03 — Control flow](../lessons/03-control-flow) | `if`, `for`, `switch`           | `for` is the only loop keyword |

After this you can: write a script that classifies inputs.

## Week 2 — Collections (~2 hours)

| Lesson | Focus | Key idea |
|---|---|---|
| [04 — Arrays](../lessons/04-arrays)     | Fixed-size, value semantics | Arrays are **value types** |
| [05 — Slices](../lessons/05-slices)     | Dynamic sequences, generics | Pointer + len + cap header |
| [06 — Maps](../lessons/06-maps)         | Hash tables                 | Iteration order is randomized on purpose |

After this you can: process word lists, build counters, return top-N.

## Week 3 — Code organization (~3 hours)

| Lesson | Focus | Key idea |
|---|---|---|
| [07 — Functions](../lessons/07-functions)   | Multiple return, variadic, closures | Errors are values |
| [08 — Methods](../lessons/08-methods)       | Receivers, interfaces               | Interfaces are **implicit** |
| [09 — Structs](../lessons/09-structs)       | Composite types, `fmt.Stringer`     | Capital letter = exported |
| [10 — Pointers](../lessons/10-pointers)     | `*T`, `&x`, `*p`, linked list       | Pass-by-value unless you take an address |

After this you can: model a domain with idiomatic types.

## Week 4 — Robustness (~2 hours)

| Lesson | Focus | Key idea |
|---|---|---|
| [11 — Errors](../lessons/11-errors)       | Sentinels, `%w`, `errors.Is`/`As` | Wrap; don't lose the chain |
| [12 — Packages](../lessons/12-packages)   | Modules, imports, exports         | One directory = one package |

After this you can: structure a real Go project.

## Week 5 — Beyond the basics (~2 hours)

| Lesson | Focus | Key idea |
|---|---|---|
| [13 — Concurrency](../lessons/13-concurrency) | Goroutines, channels, mutex, atomic | Always run with `-race` |
| [14 — Testing](../lessons/14-testing)         | Table tests, benchmarks, fuzz       | The framework ships with Go |

After this you can: ship concurrent, tested Go code.

---

## Then the challenges

Pick whichever pulls you in. Each combines several lessons.

| Challenge | Wires together |
|---|---|
| [FizzBuzz](../challenges/01-fizzbuzz)        | Control flow + `io.Writer` injection |
| [Linked List](../challenges/02-linked-list)  | Pointers + generics + in-place ops |
| [`ztail`](../challenges/03-tail)             | Filesystem + argv + exit codes |
| [Library](../challenges/04-library)          | Structs + errors + concurrency |

And the [pinball project](../projects/pinball) to see what splitting state from I/O looks like in a real program.

---

## What's next (after this course)

When the 14 lessons feel easy:

- Read **[Effective Go](https://go.dev/doc/effective_go)** end to end.
- Walk **[The Go Standard Library](https://pkg.go.dev/std)** package by package — pick five that match your work and read their docs.
- Build a small HTTP service with `net/http` + `encoding/json` + `database/sql`.
- Profile something with `pprof` and `go test -bench`.
- Read **["100 Go Mistakes and How to Avoid Them"](https://100go.co/)** to skip a year of debugging.
