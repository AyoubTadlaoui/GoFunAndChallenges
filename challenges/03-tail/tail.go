// Package main is challenge 03: ztail — a simplified `tail -c N` implementation.
//
// Usage:
//
//	go run ./challenges/03-tail -c N file [file ...]
//
// Behavior per the Zone01 spec:
//   - -c N (positive integer) is required and must come first.
//   - Prints the last N bytes of each file.
//   - When multiple files are given, prints "==> file <==\n" headers separated
//     by blank lines between files.
//   - Errors are printed inline but processing continues; exit code is 1 if
//     any file failed.
package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

// ErrUsage signals the args couldn't be parsed.
var ErrUsage = errors.New("usage: ztail -c N file [file ...]")

// parseArgs validates argv and returns the byte count and remaining file paths.
func parseArgs(argv []string) (n int, files []string, err error) {
	if len(argv) < 3 || argv[0] != "-c" {
		return 0, nil, ErrUsage
	}
	n, err = strconv.Atoi(argv[1])
	if err != nil || n < 0 {
		return 0, nil, fmt.Errorf("-c expects a non-negative integer, got %q", argv[1])
	}
	return n, argv[2:], nil
}

// tailFile writes the last n bytes of path to w. It uses Seek so it works on
// large files without reading the whole content.
func tailFile(w io.Writer, path string, n int) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}
	offset := info.Size() - int64(n)
	if offset < 0 {
		offset = 0
	}
	if _, err := f.Seek(offset, io.SeekStart); err != nil {
		return err
	}
	_, err = io.Copy(w, f)
	return err
}

// run is the testable entry point. It writes results to stdout and errors to
// stderr, and returns the process exit code.
func run(stdout, stderr io.Writer, argv []string) int {
	n, files, err := parseArgs(argv)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	multi := len(files) > 1
	exit := 0
	for i, path := range files {
		if multi {
			if i > 0 {
				fmt.Fprintln(stdout)
			}
			fmt.Fprintf(stdout, "==> %s <==\n", path)
		}
		if err := tailFile(stdout, path, n); err != nil {
			fmt.Fprintln(stderr, err)
			exit = 1
		}
	}
	return exit
}

func main() {
	os.Exit(run(os.Stdout, os.Stderr, os.Args[1:]))
}
