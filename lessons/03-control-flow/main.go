// Loops, conditions, and switch
// Exercise: Print only even numbers and add another switch case.
package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
	day := "Sunday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}
}
