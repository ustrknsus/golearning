// Channels and bounded worker pools
// Exercise: Change the worker count and add an error field to results.
package main

import (
	"fmt"
	"sync"
)

type Result struct{ Input, Square int }

func main() {
	jobs := make(chan int)
	results := make(chan Result)
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range jobs {
				results <- Result{n, n * n}
			}
		}()
	}
	go func() {
		for n := 1; n <= 8; n++ {
			jobs <- n
		}
		close(jobs)
		wg.Wait()
		close(results)
	}()
	for result := range results {
		fmt.Printf("%d squared = %d\n", result.Input, result.Square)
	}
}
