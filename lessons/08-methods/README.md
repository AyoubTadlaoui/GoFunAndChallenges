# 08 — Methods & Interfaces

Behavior attached to types. Polymorphism without inheritance.

## Run / Test

```bash
go run ./lessons/08-methods
go test ./lessons/08-methods -v
```

## What to learn

- A method is just a function with a **receiver**: `func (r Rectangle) Area() float64`.
- Use a **pointer receiver** (`*T`) when the method mutates the value or the value is large. By convention, all methods on a type pick one — don't mix.
- Interfaces are **implicit**. `Rectangle` is a `Shape` because it has `Area() float64`. No `implements` keyword, no interface registry.

## Try yourself

1. Add `Triangle` and watch `TotalArea` accept it without changing.
2. Implement `fmt.Stringer` (`String() string`) on `Rectangle` and `Circle`.
3. Re-do `Counter` as a value receiver and observe how `Inc` silently does nothing.
