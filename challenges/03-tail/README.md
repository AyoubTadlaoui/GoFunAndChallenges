# Challenge 03 — `ztail`

A simplified `tail -c N` from the Zone01 curriculum, ported here in full.

## Run

```bash
echo "abcdefghijklmnopqrstuvwxyz" > /tmp/f.txt
go run ./challenges/03-tail -c 4 /tmp/f.txt
# → xyz
```

Multiple files:

```bash
go run ./challenges/03-tail -c 4 file1.txt file2.txt
# ==> file1.txt <==
# xyz
#
# ==> file2.txt <==
# xyz
```

## Test

```bash
go test ./challenges/03-tail -v
```

## Design notes

- `Seek` straight to `size - n` instead of reading the whole file — this works on multi-gigabyte files.
- The `run` function takes `stdout` and `stderr` as parameters; that's what lets the test suite assert on the exact bytes produced. The real `main` just plumbs `os.Stdout` and `os.Stderr` through.
- Errors are reported per file but processing continues. Exit code is 1 if any file failed, 0 otherwise — matches the spec.
