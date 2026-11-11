// Command-line flags
// Exercise: Add a --verbose boolean flag.
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "learner", "name to greet")
	repeat := flag.Int("repeat", 1, "number of greetings")
	flag.Parse()
	if *repeat < 1 {
		fmt.Fprintln(os.Stderr, "repeat must be positive")
		os.Exit(1)
	}
	for i := 0; i < *repeat; i++ {
		fmt.Println("Hello,", *name)
	}
}
