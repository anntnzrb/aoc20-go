package day01

import (
	"testing"

	aocmath "aoc20-go/math"
	"aoc20-go/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSumsTo(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		target   int
		expected bool
	}{
		{
			name:     "sample numbers sum to 2020",
			numbers:  []int{1721, 299},
			target:   2020,
			expected: true,
		},
		{
			name:     "numbers don't sum to target",
			numbers:  []int{1721, 299},
			target:   1000,
			expected: false,
		},
		{
			name:     "empty slice",
			numbers:  []int{},
			target:   0,
			expected: true,
		},
		{
			name:     "single number matches target",
			numbers:  []int{2020},
			target:   2020,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := aocmath.SumsTo(tt.numbers, tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestProduct(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{
			name:     "sample product calculation",
			numbers:  []int{1721, 299},
			expected: 514579,
		},
		{
			name:     "three numbers product",
			numbers:  []int{979, 366, 675},
			expected: 241861950,
		},
		{
			name:     "empty slice returns 1",
			numbers:  []int{},
			expected: 1,
		},
		{
			name:     "single number",
			numbers:  []int{42},
			expected: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := aocmath.Product(tt.numbers)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNSum(t *testing.T) {
	sampleEntries := []int{1721, 979, 366, 299, 675, 1456}

	tests := []struct {
		name     string
		entries  []int
		target   int
		n        int
		expected []int
	}{
		{
			name:     "two sum to 2020",
			entries:  sampleEntries,
			target:   2020,
			n:        2,
			expected: []int{1721, 299},
		},
		{
			name:     "three sum to 2020",
			entries:  sampleEntries,
			target:   2020,
			n:        3,
			expected: []int{979, 366, 675},
		},
		{
			name:     "one sum (target exists)",
			entries:  sampleEntries,
			target:   1721,
			n:        1,
			expected: []int{1721},
		},
		{
			name:     "one sum (target doesn't exist)",
			entries:  sampleEntries,
			target:   9999,
			n:        1,
			expected: nil,
		},
		{
			name:     "no solution exists",
			entries:  []int{1, 2, 3},
			target:   10,
			n:        2,
			expected: nil,
		},
		{
			name:     "empty entries",
			entries:  []int{},
			target:   2020,
			n:        2,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NSum(tt.entries, tt.target, tt.n)
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
		name     string
		filename string
		expected []int
		hasError bool
	}{
		{
			name:     "sample input parsing",
			filename: "../input/day01-sample.in",
			expected: []int{1721, 979, 366, 299, 675, 1456},
			hasError: false,
		},
		{
			name:     "non-existent file",
			filename: "../input/non-existent.in",
			expected: nil,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := utils.ReadIntegers(tt.filename)
			
			if tt.hasError {
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
			filename: "../input/day01.in",
			n:        2,
			expected: 866436,
		},
		{
			name:     "part 2 - three numbers sum to 2020",
			filename: "../input/day01.in",
			n:        3,
			expected: 276650720,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entries, err := utils.ReadIntegers(tt.filename)
			require.NoError(t, err, "Failed to read input file")
			
			numbers := NSum(entries, 2020, tt.n)
			require.NotNil(t, numbers, "No solution found for %d-sum", tt.n)
			
			result := aocmath.Product(numbers)
			assert.Equal(t, tt.expected, result)
		})
	}
}