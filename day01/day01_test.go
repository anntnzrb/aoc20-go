package day01

import (
	"testing"

	"aoc20-go/aocmath"
	"aoc20-go/input"
	"aoc20-go/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSumsTo(t *testing.T) {
	type input struct {
		numbers []int
		target  int
	}

	testutil.Run(t,
		func(in input) bool { return aocmath.SumsTo(in.numbers, in.target) },
		[]struct {
			Name     string
			Input    input
			Expected bool
		}{
			testutil.T("sample numbers sum to 2020", input{[]int{1721, 299}, 2020}, true),
			testutil.T("numbers don't sum to target", input{[]int{1721, 299}, 1000}, false),
			testutil.T("empty slice", input{[]int{}, 0}, true),
			testutil.T("single number matches target", input{[]int{2020}, 2020}, true),
		})
}

func TestProduct(t *testing.T) {
	testutil.Run(t,
		func(numbers []int) int { return aocmath.Product(numbers) },
		[]struct {
			Name     string
			Input    []int
			Expected int
		}{
			testutil.T("sample product calculation", []int{1721, 299}, 514579),
			testutil.T("three numbers product", []int{979, 366, 675}, 241861950),
			testutil.T("empty slice returns 1", []int{}, 1),
			testutil.T("single number", []int{42}, 42),
		})
}

func TestNSum(t *testing.T) {
	sampleEntries := []int{1721, 979, 366, 299, 675, 1456}

	type input struct {
		entries []int
		target  int
		n       int
	}

	tests := []struct {
		name     string
		input    input
		expected []int
	}{
		{
			name:     "two sum to 2020",
			input:    input{sampleEntries, 2020, 2},
			expected: []int{1721, 299},
		},
		{
			name:     "three sum to 2020",
			input:    input{sampleEntries, 2020, 3},
			expected: []int{979, 366, 675},
		},
		{
			name:     "one sum (target exists)",
			input:    input{sampleEntries, 1721, 1},
			expected: []int{1721},
		},
		{
			name:     "one sum (target doesn't exist)",
			input:    input{sampleEntries, 9999, 1},
			expected: nil,
		},
		{
			name:     "no solution exists",
			input:    input{[]int{1, 2, 3}, 10, 2},
			expected: nil,
		},
		{
			name:     "empty entries",
			input:    input{[]int{}, 2020, 2},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NSum(Entries(tt.input.entries), Target(tt.input.target), Count(tt.input.n))
			if tt.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.Len(t, result, len(tt.expected))
				assert.ElementsMatch(t, tt.expected, result)
			}
		})
	}
}

func TestParseInput(t *testing.T) {
	tests := []struct {
		name      string
		filename  string
		expected  []int
		shouldErr bool
	}{
		{
			name:      "sample input parsing",
			filename:  "../inputs/day01-sample.in",
			expected:  []int{1721, 979, 366, 299, 675, 1456},
			shouldErr: false,
		},
		{
			name:      "non-existent file",
			filename:  "../inputs/non-existent.in",
			expected:  nil,
			shouldErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := input.ReadIntegers(tt.filename)

			if tt.shouldErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestRealInputSolutions(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		n        int
		expected int
	}{
		{
			name:     "part 1 - two numbers sum to 2020",
			filename: "../inputs/day01.in",
			n:        2,
			expected: 866436,
		},
		{
			name:     "part 2 - three numbers sum to 2020",
			filename: "../inputs/day01.in",
			n:        3,
			expected: 276650720,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entries, err := input.ReadIntegers(tt.filename)
			require.NoError(t, err, "Failed to read input file")

			combination := NSum(Entries(entries), Target(2020), Count(tt.n))
			require.NotNil(t, combination, "No combination found for %d-sum", tt.n)

			result := aocmath.Product([]int(combination))
			assert.Equal(t, tt.expected, result)
		})
	}
}
