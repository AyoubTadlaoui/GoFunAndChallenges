# 02 — Variables, Constants & Types

Declare values. Compute with them. Identify them at runtime.

## Run it

```bash
go run ./lessons/02-variables
```

## Test it

```bash
go test ./lessons/02-variables -v
```

## What to learn

- `var x int` vs `x := 0` — both declare; the short form infers the type.
- `const` values are computed at compile time. `math.Pi` is the constant you want; `math.Phi` is the **golden ratio** — a common trap fixed in this lesson.
- A type switch (`switch v.(type)`) inspects the runtime type of an `any` value.
- The 2D **area** of a disk is `πr²`. `2πr` is the **circumference**, not the "surface".

## Try yourself

1. Add `TriangleArea(base, height float64) float64` and a test.
2. Add a `complex128` case to `DescribeType`.
3. Replace `any` with a generic constraint (`[T Numeric]`) for the area helpers — see how generics tighten the API.
