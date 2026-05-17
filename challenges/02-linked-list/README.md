# Challenge 02 — Doubly Linked List

A generic doubly linked list — `List[T]` for any T — with O(1) insert at both ends, O(1) removal at any node, and in-place reverse.

## Run / Test

```bash
go run ./challenges/02-linked-list
go test ./challenges/02-linked-list -v
```

## Where this differs from lesson 10

- **Doubly linked**: each node knows both neighbors, so `Remove(node)` is O(1) without searching.
- **Generic**: `List[T]` works for any type.
- **Reverse in place** flips every `prev`/`next` pointer and swaps head/tail. No allocations.

## Extend it

- Add `MoveToFront(n *Node[T])`.
- Wrap it as an LRU cache: `Get`/`Put` with a `map[K]*Node[entry]` for O(1) lookup.
- Make it concurrent-safe with a `sync.RWMutex`.
