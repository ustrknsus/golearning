// Package greeting demonstrates a reusable internal package.
package greeting

import "fmt"

func Hello(name string) string { return fmt.Sprintf("Hello, %s!", name) }
