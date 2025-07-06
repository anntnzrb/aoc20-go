package testutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestUtils(t *testing.T) {
	// Test the T helper function
	testCase := T("test case", 42, "expected")
	assert.Equal(t, "test case", testCase.Name)
	assert.Equal(t, 42, testCase.Input)
	assert.Equal(t, "expected", testCase.Expected)

	// Test the Run function
	Run(t,
		func(x int) int { return x * 2 },
		[]struct {
			Name     string
			Input    int
			Expected int
		}{
			T("double of 5", 5, 10),
			T("double of 0", 0, 0),
			T("double of negative", -3, -6),
		})
}
