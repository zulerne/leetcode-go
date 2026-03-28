# LeetCode Solutions in Go

## Workflow

```bash
# Scaffold a new problem
task new -- 0042-trapping-rain-water

# Run tests
task run-v -- 0042-trapping-rain-water

# Benchmark
task bench -- 0042-trapping-rain-water

# Format code
task fmt
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

`task new -- 0009-palindrome-number` scaffolds both files with the correct package name and LeetCode URL. Just write the function signature and tests.
