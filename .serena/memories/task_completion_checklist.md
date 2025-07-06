# Task Completion Checklist

## When Completing Any Task

### 1. Testing
- [ ] Run all tests: `go test ./...`
- [ ] Ensure all tests pass
- [ ] Add tests for new functionality
- [ ] Verify test coverage is adequate

### 2. Code Quality
- [ ] Format code: `gofmt -w .`
- [ ] Run static analysis: `go vet ./...`
- [ ] Check for any linting issues
- [ ] Ensure code follows project conventions

### 3. TDD Specific
- [ ] Follow Red-Green-Refactor cycle
- [ ] Write failing tests first (Red)
- [ ] Implement minimal code to pass (Green)
- [ ] Refactor while maintaining green tests
- [ ] Ensure components test core business logic

## Before Final Submission
- [ ] All tests pass
- [ ] Code is properly formatted
- [ ] No vet warnings
- [ ] Changes are committed
