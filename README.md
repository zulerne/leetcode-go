# LeetCode Solutions in Go

## Workflow

```bash
# Scaffold a new problem (number is auto-padded: 42 → 0042)
task new -- 42-trapping-rain-water

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

`task new -- 9-palindrome-number` scaffolds both files with the correct package name and LeetCode URL (auto-pads to `0009-palindrome-number`). Just write the function signature and tests.
