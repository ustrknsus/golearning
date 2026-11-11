// External commands with deadlines
// Exercise: Capture stdout and stderr separately; avoid sh -c for untrusted input.
package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	output, err := exec.CommandContext(ctx, "go", "version").CombinedOutput()
	if err != nil {
		return fmt.Errorf("run command: %w; output: %s", err, output)
	}
	fmt.Print(string(output))
	return nil
}
func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
