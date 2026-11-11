// Strings, runes, and Unicode
// Exercise: Compare len with utf8.RuneCountInString for your own text.
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	text := "Olá, Go"
	fmt.Println("Bytes:", len(text), "Runes:", utf8.RuneCountInString(text))
	fmt.Println(strings.ToUpper(text))
	for index, character := range text {
		fmt.Printf("byte offset %d: %c\n", index, character)
	}
}
