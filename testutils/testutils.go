// Package testutils provides utilities for reducing boilerplate in table-driven tests.
//
// This package offers simple helpers to create and execute table-driven tests
// without having to manually define struct types and iteration logic.
//
// Example usage:
//
//	func TestAdd(t *testing.T) {
//		type input struct {
//			a, b int
//		}
//
//		testutils.Run(t,
//			func(in input) int { return in.a + in.b },
//			[]struct {
//				Name     string
//				Input    input
//				Expected int
//			}{
//				testutils.T("simple addition", input{2, 3}, 5),
//				testutils.T("zero addition", input{0, 5}, 5),
//				testutils.T("negative numbers", input{-1, 1}, 0),
//			})
//	}
package testutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// T creates a test case struct with the given name, input, and expected output.
// This eliminates the need to manually construct test case structs in table-driven tests.
//
// Example:
//
//	testutils.T("test name", inputValue, expectedOutput)
//
// Returns a struct with Name, Input, and Expected fields that can be used
// directly in test case slices.
func T[I, O any](name string, input I, expected O) struct {
	Name     string
	Input    I
	Expected O
} {
	return struct {
		Name     string
		Input    I
		Expected O
	}{name, input, expected}
}

// Run executes table-driven tests for functions with a single input and output.
// It automatically iterates through the test cases, runs each as a subtest,
// and asserts that the function output equals the expected result.
//
// Parameters:
//   - t: The testing.T instance
//   - fn: The function to test (must take input of type I and return output of type O)
//   - cases: Slice of test cases with Name, Input, and Expected fields
//
// Example:
//
//	testutils.Run(t,
//		func(x int) int { return x * 2 },
//		[]struct {
//			Name     string
//			Input    int
//			Expected int
//		}{
//			testutils.T("double of 5", 5, 10),
//			testutils.T("double of 0", 0, 0),
//		})
//
// Each test case will be executed as a subtest using t.Run() with the provided name.
// The function result will be compared to the expected value using assert.Equal().
func Run[I, O any](t *testing.T, fn func(I) O, cases []struct {
	Name     string
	Input    I
	Expected O
}) {
	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			result := fn(tt.Input)
			assert.Equal(t, tt.Expected, result)
		})
	}
}