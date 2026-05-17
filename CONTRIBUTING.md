# Contributing

Thanks for considering a contribution. This repo is a teaching codebase — clarity beats cleverness everywhere.

## Workflow

1. **Fork & branch** off `main`. Use a short topical name (`fix-leap-year`, `add-lesson-15-generics`).
2. **Run the sanity check** locally before opening a PR:
   ```bash
   make check        # gofmt + go vet + go test
   make test-race    # if you touched concurrent code
   ```
3. **Open a PR** with:
   - One sentence explaining *what* and *why*.
   - Test output (or screenshots for any UX change).
   - A note on any docs you updated.

## What we love in a PR

- A new lesson or challenge that follows the existing shape (`main.go` + `<topic>.go` + `<topic>_test.go` + `README.md`).
- A bug fix paired with a regression test.
- README / docs improvements that lower the bar for a beginner.
- Replacing a hand-rolled helper with a standard-library one (with the rationale in the commit message).

## What slows a PR down

- Adding a dependency without a clear reason. Stdlib first.
- Reformatting unrelated files in the same diff.
- Lessons without tests, or tests without assertions.
- "Refactor for elegance" without a measurable benefit (readability counts as one, but say so).

## Style

- `make fmt` and `make vet` must pass.
- `golangci-lint run` (`make lint`) should not introduce new warnings.
- Comments explain *why*, not *what*. Well-named identifiers do the *what*.
- Each exported symbol has a doc comment starting with its name (the Go convention).

## Adding a new lesson

```text
lessons/
└── NN-topic/
    ├── README.md          # goal, run command, key takeaways, "Try yourself"
    ├── main.go            # tiny demo (package main, func main)
    ├── topic.go           # the library code (same package main)
    └── topic_test.go      # table-driven tests
```

Update the lesson table in [`README.md`](README.md) and the path in [`docs/CURRICULUM.md`](docs/CURRICULUM.md).

## Reporting issues

Open one with:

- The lesson / challenge / file in question.
- What you ran and what you expected vs. what happened.
- Your `go version` output.
