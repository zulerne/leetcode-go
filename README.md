# LeetCode Solutions in Go

## Workflow

```bash
# Scaffold a new problem
make new PKG=0042-trapping-rain-water

# Run tests
make run-v PKG=0042-trapping-rain-water

# Benchmark
make bench PKG=0042-trapping-rain-water

# Format code
make fmt
```

## Structure

```
problems/
  NNNN-slug/
    solution.go       # solution
    solution_test.go  # tests + benchmark
kit/
  assert.go           # test helpers
```

### kit helpers

```go
kit.AssertEqual(t, got, want)
kit.AssertTrue(t, condition)
kit.AssertFalse(t, condition)
kit.AssertNil(t, value)
kit.AssertNotNil(t, value)
```

## Adding a new problem

`make new PKG=0009-palindrome-number` scaffolds both files with the correct package name and LeetCode URL. Just write the function signature and tests.
