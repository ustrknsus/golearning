// Third-party modules and dependency management
// Exercise: Inspect go.mod and go.sum, then run go list -m all and go mod graph.
package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() { fmt.Println("Generated UUID:", uuid.NewString()) }
