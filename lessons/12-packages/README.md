# 12 — Packages & Modules

A **package** is the unit of code (one directory). A **module** is the unit of versioning (one `go.mod`).

## Run / Test

```bash
go run ./lessons/12-packages
go test ./lessons/12-packages/...
```

## What to learn

- Each directory is exactly one package. The package name (here `calc`) usually matches the directory name but doesn't have to.
- An **import path** is the module path plus the relative directory: `github.com/AyoubTadlaoui/GoFunAndChallenges/lessons/12-packages/calc`.
- Only **uppercase** identifiers are visible outside the package (`Add`, `ErrDivByZero`). Lowercase stays private.

## Try yourself

1. Add a `geo` sub-package with `Rect` and `Circle` types and use them from `main.go`.
2. Add a `_test.go` in a separate `calc_test` package (black-box testing) and verify only exported names are accessible.
