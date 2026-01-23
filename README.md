# LeetCode Solutions in Go

A collection of LeetCode problem solutions implemented in Go, focusing on clean code, performance optimization, and
comprehensive test coverage.

## 📁 Project Structure

```
leetcode-go/
├── kit/              # Testing utilities
│   └── assert.go     # Custom assertion helper
└── problems/         # LeetCode solutions
    ├── 0000-template/
    ├── 0001-two-sum/
    ├── 0011-container-with-most-water/
    └── ...
```

Each problem directory contains:

- `solution.go` - Implementation with complexity analysis
- `solution_test.go` - Comprehensive test cases

## 🚀 Getting Started

### Using Makefile

```bash
# Run all tests
make test

# Run tests with verbose output
make test-v

# Generate coverage report
make test-cover

# Format code
make fmt

# Create new problem from template
make new NUM=0042 NAME=trapping-rain-water
```

### Running Tests Directly

```bash
# Run tests for a specific problem
go test ./problems/0001-two-sum

# Run tests with verbose output
go test -v ./problems/0001-two-sum
```

## 📚 Template

Use the template in `problems/0000-template` to add new solutions:

```go
// Package {template}
// {url}
package template

func function(arg string) string {
	return ""
}
```

## 🧪 Testing Utilities

The `kit` package provides assertion helpers:

```go
kit.AssertEqual(t, got, want)     // Deep equality check
kit.AssertTrue(t, condition)       // Check condition is true
kit.AssertFalse(t, condition)      // Check condition is false
kit.AssertNil(t, value)            // Check value is nil
kit.AssertNotNil(t, value)         // Check value is not nil
```

All helpers use `t.Helper()` for clear error messages pointing to test cases.

## 📄 License

This project is open source and available for educational purposes.

---

**Note**: Problem links reference leetcode.com. Solutions are implemented for learning and practice purposes.