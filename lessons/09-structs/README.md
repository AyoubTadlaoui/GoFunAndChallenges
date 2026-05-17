# 09 — Structs

Group related fields. Nest them. Attach behavior.

## Run / Test

```bash
go run ./lessons/09-structs
go test ./lessons/09-structs -v
```

## What to learn

- Field names that start uppercase are **exported** (visible to other packages); lowercase are package-private.
- Implementing `String() string` makes `fmt.Println(p)` use your formatting (the `fmt.Stringer` interface).
- A value receiver means the method gets a **copy** — that's how `WithEmail` returns an updated copy without touching the original.

## Try yourself

1. Add a `Phone` field as `*string` and use it to model "phone optional, may be nil".
2. Implement `(p Person) MarshalJSON() ([]byte, error)` to emit only `name` and `country`.
3. Write a `ByAge` type that implements `sort.Interface`, then sort `[]Person` with it.
