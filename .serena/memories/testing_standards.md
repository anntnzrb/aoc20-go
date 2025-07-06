# Testing Standards

## Test Structure
- **Use table-driven tests** for all test functions
- Group related test cases into logical test functions
- Use descriptive test names that clearly indicate what is being tested

## Table-Driven Test Pattern
```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name     string
        input    InputType
        expected OutputType
        // additional fields as needed
    }{
        {
            name:     "descriptive test case name",
            input:    someInput,
            expected: expectedOutput,
        },
        // more test cases...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := FunctionUnderTest(tt.input)
            assert.Equal(t, tt.expected, result)
        })
    }
}
```

## Test Coverage
- Aim for 100% statement coverage
- Include edge cases: empty inputs, nil values, error conditions
- Test both success and failure scenarios
- Use `testify/assert` and `testify/require` for assertions

## Test Organization
- Group related functionality into single test functions
- Use clear, descriptive test case names
- Keep test cases focused and atomic
- Avoid performance benchmarks unless specifically needed
