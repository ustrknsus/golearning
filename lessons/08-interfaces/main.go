// Interfaces and dependency injection
// Exercise: Implement a second notifier that writes to a file.
package main

import "fmt"

type Notifier interface{ Notify(string) error }
type ConsoleNotifier struct{}

func (ConsoleNotifier) Notify(message string) error { fmt.Println("Notice:", message); return nil }
func announce(n Notifier) error                     { return n.Notify("Learning interfaces") }
func main() {
	if err := announce(ConsoleNotifier{}); err != nil {
		fmt.Println(err)
	}
}
