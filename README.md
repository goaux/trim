# trim

[![Go Reference](https://pkg.go.dev/badge/github.com/goaux/trim.svg)](https://pkg.go.dev/github.com/goaux/trim)
[![Go Report Card](https://goreportcard.com/badge/github.com/goaux/trim)](https://goreportcard.com/report/github.com/goaux/trim)

The `trim` package provides functions to handle indentation and whitespace in multiline strings.

This package is inspired by similar utility functions in the Kotlin standard library.

## Features

- `Indent()`: Removes a common minimal indent from all input lines
- `Margin()`: Removes a common prefix from each line
- `IsBlank()`: Checks if a string contains only whitespace characters

## Installation

```bash
go get github.com/goaux/trim
```

## Usage Examples

### Indent

```go
text := `
    ABC
      123
    456
`
result := trim.Indent(text)
// result: "ABC\n  123\n456"
```

### Margin

```go
text := `
    > ABC
    >   123
    > 456
`
result := trim.Margin(text, "> ")
// result: "ABC\n  123\n456"
```

### IsBlank

```go
blank := trim.IsBlank("   \t\n")  // true
notBlank := trim.IsBlank("hello") // false
```
