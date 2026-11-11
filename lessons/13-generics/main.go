// Generic functions
// Exercise: Implement a generic Contains function.
package main

import "fmt"

func Map[T any, U any](items []T, transform func(T) U) []U {
	result := make([]U, len(items))
	for i, item := range items {
		result[i] = transform(item)
	}
	return result
}
func main() {
	squares := Map([]int{1, 2, 3, 4}, func(n int) int { return n * n })
	fmt.Println(squares)
}
