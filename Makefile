.PHONY: run run-v bench test-all test-cover new fmt

# ============================================
# SINGLE PROBLEM WORKFLOW (main commands)
# ============================================

# Run tests for a specific problem
# Usage: make run PKG=0001-two-sum
run:
	@if [ -z "$(PKG)" ]; then \
		echo "Usage: make run PKG=0001-two-sum"; \
		exit 1; \
	fi
	go test ./problems/$(PKG)

# Run tests with verbose output for a specific problem
# Usage: make run-v PKG=0001-two-sum
run-v:
	@if [ -z "$(PKG)" ]; then \
		echo "Usage: make run-v PKG=0001-two-sum"; \
		exit 1; \
	fi
	go test -v ./problems/$(PKG)

# Run benchmarks for a specific problem
# Usage: make bench PKG=0001-two-sum
bench:
	@if [ -z "$(PKG)" ]; then \
		echo "Usage: make bench PKG=0001-two-sum"; \
		exit 1; \
	fi
	go test -bench=. -benchmem ./problems/$(PKG)

# ============================================
# ALL PROBLEMS (CI / full check)
# ============================================

# Run all tests
test-all:
	go test ./...

# Run tests with coverage
test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

# ============================================
# UTILITIES
# ============================================

# Format code
fmt:
	go fmt ./...

# Create new problem directory
# Usage: make new PKG=0042-trapping-rain-water
new:
	@if [ -z "$(PKG)" ]; then \
		echo "Usage: make new PKG=0042-problem-name"; \
		exit 1; \
	fi
	@mkdir -p problems/$(PKG)
	@cp problems/0000-template/solution.go problems/$(PKG)/
	@cp problems/0000-template/solution_test.go problems/$(PKG)/
	@sed -i '' 's/\/\/go:build ignore//' problems/$(PKG)/solution.go
	@sed -i '' 's/\/\/go:build ignore//' problems/$(PKG)/solution_test.go
	@echo "Created: problems/$(PKG)/"
	@echo "Run: make run PKG=$(PKG)"


