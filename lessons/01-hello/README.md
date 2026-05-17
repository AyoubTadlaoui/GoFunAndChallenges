# 01 — Hello, Go

Your first runnable Go program.

## Run it

```bash
go run ./lessons/01-hello
```

You should see:

```
Hello, Gopher! Welcome to Go.
```

## Test it

```bash
go test ./lessons/01-hello
```

## What to learn

- `package main` + `func main()` make a runnable binary.
- `fmt.Println` writes to standard output.
- Pure functions like `Greet` are easy to test — that's why the printing lives in `main` and the logic lives in `Greet`.

## Try yourself

1. Make `Greet` accept a second argument `language` (e.g. `"fr"`, `"ar"`, `"en"`) and return the right greeting.
2. Add a test case for each language.
3. Run `go test ./lessons/01-hello -v` and watch all cases pass.
