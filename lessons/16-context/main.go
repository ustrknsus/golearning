// Context, select, and cancellation
// Exercise: Replace the deadline with explicit cancellation.
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 350*time.Millisecond)
	defer cancel()
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopped:", ctx.Err())
			return
		case <-ticker.C:
			fmt.Println("Working")
		}
	}
}
