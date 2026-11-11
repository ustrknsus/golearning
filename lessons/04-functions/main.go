// Functions, multiple returns, and closures
// Exercise: Write a variadic sum function.
package main

import "fmt"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
func counter() func() int {
	n := 0
	return func() int { n++; return n }
}
func main() {
	result, err := divide(12, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Result:", result)
	next := counter()
	fmt.Println("Closure:", next(), next())
}
