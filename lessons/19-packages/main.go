// Modules and internal packages
// Exercise: Add another exported function in internal/greeting.
package main

import (
	"example.com/go-learning/internal/greeting"
	"fmt"
)

func main() { fmt.Println(greeting.Hello("Go learner")) }
