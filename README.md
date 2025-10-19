# Go Getting Started Tutorial

As part of my onboarding to the Bulls team - which is responsble on a service written in Go.

This repository contains a Go implementation following the official [Go Getting Started Tutorial](https://go.dev/doc/tutorial/getting-started). The project demonstrates fundamental Go concepts including modules, packages, external dependencies, and basic programming patterns.

## Project Structure

```
go_getting_started/
├── greetings/           # Custom greetings module
│   ├── go.mod          # Module definition
│   ├── greetings.go    # Main greetings package
│   └── greetings_test.go # Unit tests
├── hello/              # Main application
│   ├── go.mod          # Module definition with dependencies
│   ├── go.sum          # Dependency checksums
│   └── hello.go        # Main application code
└── README.md           # This file
```

## What This Tutorial Covers

This project implements the following Go concepts from the official tutorial:

1. **Module Creation**: Creating Go modules with `go mod init`
2. **Package Management**: Organizing code into packages
3. **External Dependencies**: Using third-party packages from pkg.go.dev
4. **Function Implementation**: Creating and calling functions
5. **Error Handling**: Proper Go error handling patterns
6. **Testing**: Writing and running unit tests

## Prerequisites

- Go 1.25.1 or later installed on your system
- A text editor (VSCode, GoLand, Vim, etc.)
- Command terminal access

## Getting Started

### 1. Clone and Navigate

```bash
git clone <your-repo-url>
cd go_getting_started
```

### 2. Explore the Greetings Module

The `greetings/` directory contains a custom module that demonstrates:

- **Function Creation**: `Hello()` function with error handling
- **Random Messages**: Dynamic greeting generation
- **Multiple Names**: `Hellos()` function for handling multiple names
- **Error Handling**: Proper Go error patterns

```bash
cd greetings
go run greetings.go
```

### 3. Run the Main Application

The `hello/` directory contains the main application that:

- Uses the custom `greetings` module
- Imports external packages (`rsc.io/quote`)
- Demonstrates module dependencies

```bash
cd hello
go run hello.go
```

Expected output:
```
Hello, World from Brosh!
Don't communicate by sharing memory, share memory by communicating.
Hi, Nevo. Welcome from Brosh!
```

## Key Features Demonstrated

### Custom Greetings Module

The `greetings` package provides:

- **`Hello(name string)`**: Returns a personalized greeting with error handling
- **`Hellos(names []string)`**: Handles multiple names and returns a map
- **`randomFormat()`**: Generates random greeting formats

### External Package Usage

The project uses `rsc.io/quote` package to demonstrate:

- Finding packages on [pkg.go.dev](https://pkg.go.dev)
- Adding external dependencies with `go mod tidy`
- Calling functions from external modules

### Error Handling

Proper Go error handling patterns:

```go
msg, err := greetings.Hello("Nevo")
if err != nil {
    log.Fatal(err)
}
```

## Running Tests

The project includes unit tests for the greetings module:

```bash
cd greetings
go test
```

To run tests with verbose output:

```bash
go test -v
```

## Module Dependencies

The project demonstrates Go module management:

- **Local Module**: `github.com/yuvbro/go-greetings` (custom greetings package)
- **External Module**: `rsc.io/quote v1.5.2` (third-party package)
- **Replace Directive**: Local module replacement for development

## Key Go Concepts Learned

1. **Modules**: Go's dependency management system
2. **Packages**: Code organization and reusability
3. **Imports**: Using standard library and external packages
4. **Functions**: Creating and calling functions with parameters and return values
5. **Error Handling**: Go's explicit error handling approach
6. **Testing**: Writing and running unit tests
7. **go mod**: Managing dependencies and module versions

## Next Steps

After completing this tutorial, consider exploring:

- [Create a Go Module](https://go.dev/doc/tutorial/create-module) - Next tutorial in the series
- [Go Tour](https://tour.golang.org/) - Interactive Go language tour
- [Effective Go](https://golang.org/doc/effective_go.html) - Go best practices
- [Go by Example](https://gobyexample.com/) - Hands-on Go examples

## Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go Getting Started Tutorial](https://go.dev/doc/tutorial/getting-started)
- [Go Package Discovery](https://pkg.go.dev/)
- [Go Modules Reference](https://golang.org/ref/mod)

## License

This project is for educational purposes following the Go getting started tutorial.
