// Variables, constants, and types
// Exercise: Add a float64 variable and print its type with %T.
package main

import "fmt"

func main() {
	const language = "Go"
	name := "ustrknsus"
	var lessons int = 20
	learning := true
	fmt.Printf("%s is learning %s: %d lessons, active=%t\n", name, language, lessons, learning)
}
