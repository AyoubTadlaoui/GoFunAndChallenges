# 10 — Pointers & Linked List

`*T` is a pointer to T. `&x` takes the address. `*p` dereferences. That's most of it.

## Run / Test

```bash
go run ./lessons/10-pointers
go test ./lessons/10-pointers -v
```

## What to learn

- Without pointers, a function can't mutate the caller's value — Go passes by value.
- The previous version of this list called its insert function `append` but it actually **prepended** (O(1) at the head). This version splits them: `Append` (tail) and `Prepend` (head), both O(1) because we track the tail.
- A method on `*LinkedList` can mutate the list. A method on `LinkedList` (value) would not.

## Try yourself

1. Add `Reverse()` that flips the links in place — no allocations.
2. Add `Remove(v int) bool` that drops the first matching node.
3. Detect cycles with the tortoise-and-hare technique.
