# Learning Go: Initial Lessons

This new repository is a hands-on learning project for studying Go, from
Hello, World to intermediate development, concurrency, and external modules.
Examples are runnable starting points, not production applications or evidence
of completed coursework. Read each lesson, change it, and complete its exercise.

## Requirements

- Go 1.22 or newer and Git.
- Git user.name and user.email configured for the initial commit.
- Internet access for the optional external dependency lesson.

## Start here

```bash
go run ./lessons/01-hello
go run ./lessons/11-cli --name ustrknsus --repeat 3
go test ./...
go test -race ./...
go vet ./...
go test -bench=. ./lessons/12-testing
go build -o bin/hello ./lessons/01-hello
```

Each lesson is its own package/directory, avoiding conflicting main functions.

## Learning path

| Directory | Topic |
|---|---|
| 01-hello | Hello, world |
| 02-variables | Variables, constants, and types |
| 03-control-flow | Loops, conditions, and switch |
| 04-functions | Functions, multiple returns, and closures |
| 05-collections | Arrays, slices, and maps |
| 06-strings | Strings, runes, and Unicode |
| 07-structs | Structs, pointers, and methods |
| 08-interfaces | Interfaces and dependency injection |
| 09-errors | Errors, wrapping, and defer |
| 10-files-json | Files and JSON |
| 11-cli | Command-line flags |
| 12-testing | Unit tests and benchmarks |
| 13-generics | Generic functions |
| 14-goroutines | Goroutines and synchronization |
| 15-worker-pool | Channels and bounded worker pools |
| 16-context | Context, select, and cancellation |
| 17-http | HTTP clients, servers, and timeouts |
| 18-processes | External commands with deadlines |
| 19-packages | Modules and internal packages |
| 20-external-module | Third-party modules and dependency management |

## External module lesson

Lesson 20 is deliberately a separate nested module so the initial lessons
can be built and tested without downloading external dependencies.

```bash
cd lessons/20-external-module
go get github.com/google/uuid@v1.6.0
go mod tidy
go run .
go list -m all
go mod graph
go mod verify
```

Commit go.mod and go.sum after adding a dependency. Root-level `go test ./...`
does not include this nested module. Run its checks from its own directory.
The dependency version above is pinned for reproducibility, not a claim that
it is the newest version. Review dependencies before using them in production.

## Next steps

- Add table-driven tests and error paths to every lesson.
- Combine the worker pool, context, and process lessons into a parallel command runner.
- Explore `go work init . ./lessons/20-external-module` for a multi-module workspace.
- Later, add SSH with key authentication and host-key verification.

## References

- https://go.dev/tour/
- https://go.dev/doc/tutorial/
- https://go.dev/ref/mod
