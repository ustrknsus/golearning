// Errors, wrapping, and defer
// Exercise: Wrap another error and detect it using errors.Is.
package main

import (
	"errors"
	"fmt"
)

var ErrMissing = errors.New("item missing")

func lookup() error { return fmt.Errorf("lookup server: %w", ErrMissing) }
func main() {
	defer fmt.Println("Cleanup runs when main returns")
	err := lookup()
	if errors.Is(err, ErrMissing) {
		fmt.Println("Recognized wrapped error:", err)
	}
}
