// Goroutines and synchronization
// Exercise: Launch ten tasks and observe completion order.
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) { defer wg.Done(); fmt.Println("Task", id) }(i)
	}
	wg.Wait()
}
