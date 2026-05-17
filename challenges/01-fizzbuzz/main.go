package main

import "os"

func main() {
	if err := FizzBuzz(os.Stdout, 30); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}
