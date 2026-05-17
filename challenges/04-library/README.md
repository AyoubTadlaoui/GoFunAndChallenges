# Challenge 04 — Library Manager

A small, concurrency-safe catalog API: add books, lend, return, query by author, see what's on the shelf.

## Run / Test

```bash
go run ./challenges/04-library
go test ./challenges/04-library -race -v
```

## What it shows

- **Sentinel errors** + `%w` wrapping for caller-friendly inspection (`errors.Is`).
- A `sync.RWMutex` so reads (`Get`, `ByAuthor`, `AvailableTitles`) don't block each other.
- A concurrency stress test using `sync.WaitGroup` — pair it with `-race` to prove safety.

## Extend it

- Add `Search(query string) []Book` with simple substring matching.
- Persist the catalog to JSON on disk and re-load on startup.
- Expose it over HTTP with `net/http` — instant library REST API.
