package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
}

func TestParseArgs(t *testing.T) {
	if _, _, err := parseArgs([]string{"-c", "10", "a.txt"}); err != nil {
		t.Fatalf("valid args returned err: %v", err)
	}
	if _, _, err := parseArgs([]string{"-c", "-1", "a.txt"}); err == nil {
		t.Fatal("negative n should error")
	}
	if _, _, err := parseArgs([]string{"-c", "abc", "a.txt"}); err == nil {
		t.Fatal("non-numeric n should error")
	}
	if _, _, err := parseArgs(nil); err == nil {
		t.Fatal("empty args should error")
	}
}

func TestRun_SingleFile(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "f.txt")
	writeFile(t, p, "abcdefghijklmnopqrstuvwxyz\n")

	var stdout, stderr bytes.Buffer
	exit := run(&stdout, &stderr, []string{"-c", "4", p})
	if exit != 0 {
		t.Fatalf("exit = %d (stderr=%q)", exit, stderr.String())
	}
	if got, want := stdout.String(), "xyz\n"; got != want {
		t.Fatalf("stdout = %q, want %q", got, want)
	}
}

func TestRun_MultipleFiles(t *testing.T) {
	dir := t.TempDir()
	a := filepath.Join(dir, "a.txt")
	b := filepath.Join(dir, "b.txt")
	writeFile(t, a, "abcdefghijklmnopqrstuvwxyz\n")
	writeFile(t, b, "abcdefghijklmnopqrstuvwxyz\n")

	var stdout, stderr bytes.Buffer
	exit := run(&stdout, &stderr, []string{"-c", "4", a, b})
	if exit != 0 {
		t.Fatalf("exit = %d (stderr=%q)", exit, stderr.String())
	}
	want := "==> " + a + " <==\nxyz\n\n==> " + b + " <==\nxyz\n"
	if stdout.String() != want {
		t.Fatalf("stdout =\n%q\nwant\n%q", stdout.String(), want)
	}
}

func TestRun_MissingFileExitNonZero(t *testing.T) {
	dir := t.TempDir()
	good := filepath.Join(dir, "good.txt")
	writeFile(t, good, "abcdefghijklmnopqrstuvwxyz\n")
	missing := filepath.Join(dir, "missing.txt")

	var stdout, stderr bytes.Buffer
	exit := run(&stdout, &stderr, []string{"-c", "4", good, missing})
	if exit == 0 {
		t.Fatal("exit = 0, want non-zero when a file is missing")
	}
	if stderr.Len() == 0 {
		t.Fatal("expected stderr message for missing file")
	}
}

func TestRun_NLargerThanFile(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "small.txt")
	writeFile(t, p, "hi\n")
	var stdout, stderr bytes.Buffer
	exit := run(&stdout, &stderr, []string{"-c", "1000", p})
	if exit != 0 {
		t.Fatalf("exit = %d (stderr=%q)", exit, stderr.String())
	}
	if got, want := stdout.String(), "hi\n"; got != want {
		t.Fatalf("stdout = %q, want %q", got, want)
	}
}
