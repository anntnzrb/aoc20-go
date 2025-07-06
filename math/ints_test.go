package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumsTo(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		target   int
		expected bool
	}{
		{
			name:     "basic sum matches",
			numbers:  []int{1, 2, 3},
			target:   6,
			expected: true,
		},
		{
			name:     "basic sum doesn't match",
			numbers:  []int{1, 2, 3},
			target:   7,
			expected: false,
		},
		{
			name:     "empty slice",
			numbers:  []int{},
			target:   0,
			expected: true,
		},
		{
			name:     "empty slice non-zero target",
			numbers:  []int{},
			target:   5,
			expected: false,
		},
		{
			name:     "single element matches",
			numbers:  []int{42},
			target:   42,
			expected: true,
		},
		{
			name:     "single element doesn't match",
			numbers:  []int{42},
			target:   43,
			expected: false,
		},
		{
			name:     "negative numbers",
			numbers:  []int{-1, -2, -3},
			target:   -6,
			expected: true,
		},
		{
			name:     "mixed positive and negative",
			numbers:  []int{10, -5, 2, -3},
			target:   4,
			expected: true,
		},
		{
			name:     "zero sum",
			numbers:  []int{5, -5, 3, -3},
			target:   0,
			expected: true,
		},
		{
			name:     "large numbers",
			numbers:  []int{1000000, 2000000, 3000000},
			target:   6000000,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumsTo(tt.numbers, tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSumsToFloat64(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []float64
		target   float64
		expected bool
	}{
		{
			name:     "basic float sum",
			numbers:  []float64{1.5, 2.5, 3.0},
			target:   7.0,
			expected: true,
		},
		{
			name:     "decimal precision",
			numbers:  []float64{0.25, 0.25, 0.5},
			target:   1.0,
			expected: true,
		},
		{
			name:     "negative floats",
			numbers:  []float64{-1.5, -2.5},
			target:   -4.0,
			expected: true,
		},
		{
			name:     "mixed float values",
			numbers:  []float64{1.1, -0.1, 2.0},
			target:   3.0,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumsTo(tt.numbers, tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSumsToUint(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []uint
		target   uint
		expected bool
	}{
		{
			name:     "basic uint sum",
			numbers:  []uint{1, 2, 3},
			target:   6,
			expected: true,
		},
		{
			name:     "large uint values",
			numbers:  []uint{4294967295, 1},
			target:   4294967296,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumsTo(tt.numbers, tt.target)
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
			name:     "basic product",
			numbers:  []int{2, 3, 4},
			expected: 24,
		},
		{
			name:     "empty slice",
			numbers:  []int{},
			expected: 1,
		},
		{
			name:     "single element",
			numbers:  []int{42},
			expected: 42,
		},
		{
			name:     "with zero",
			numbers:  []int{1, 2, 0, 4},
			expected: 0,
		},
		{
			name:     "negative numbers even count",
			numbers:  []int{-2, -3},
			expected: 6,
		},
		{
			name:     "negative numbers odd count",
			numbers:  []int{-2, -3, -1},
			expected: -6,
		},
		{
			name:     "mixed positive and negative",
			numbers:  []int{2, -3, 4},
			expected: -24,
		},
		{
			name:     "ones",
			numbers:  []int{1, 1, 1, 1},
			expected: 1,
		},
		{
			name:     "large numbers",
			numbers:  []int{100, 200},
			expected: 20000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Product(tt.numbers)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestProductFloat64(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []float64
		expected float64
	}{
		{
			name:     "basic float product",
			numbers:  []float64{2.5, 4.0},
			expected: 10.0,
		},
		{
			name:     "decimal values",
			numbers:  []float64{0.5, 0.2, 10.0},
			expected: 1.0,
		},
		{
			name:     "negative floats",
			numbers:  []float64{-2.0, 3.0},
			expected: -6.0,
		},
		{
			name:     "with zero",
			numbers:  []float64{1.5, 0.0, 2.0},
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Product(tt.numbers)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestProductUint(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []uint
		expected uint
	}{
		{
			name:     "basic uint product",
			numbers:  []uint{2, 3, 4},
			expected: 24,
		},
		{
			name:     "large uint values",
			numbers:  []uint{1000, 1000},
			expected: 1000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Product(tt.numbers)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestNumericConstraintTypes(t *testing.T) {
	// Test that all numeric types work with our constraint
	t.Run("int8", func(t *testing.T) {
		assert.True(t, SumsTo([]int8{1, 2, 3}, int8(6)))
		assert.Equal(t, int8(6), Product([]int8{1, 2, 3}))
	})

	t.Run("int16", func(t *testing.T) {
		assert.True(t, SumsTo([]int16{1, 2, 3}, int16(6)))
		assert.Equal(t, int16(6), Product([]int16{1, 2, 3}))
	})

	t.Run("int32", func(t *testing.T) {
		assert.True(t, SumsTo([]int32{1, 2, 3}, int32(6)))
		assert.Equal(t, int32(6), Product([]int32{1, 2, 3}))
	})

	t.Run("int64", func(t *testing.T) {
		assert.True(t, SumsTo([]int64{1, 2, 3}, int64(6)))
		assert.Equal(t, int64(6), Product([]int64{1, 2, 3}))
	})

	t.Run("uint8", func(t *testing.T) {
		assert.True(t, SumsTo([]uint8{1, 2, 3}, uint8(6)))
		assert.Equal(t, uint8(6), Product([]uint8{1, 2, 3}))
	})

	t.Run("uint16", func(t *testing.T) {
		assert.True(t, SumsTo([]uint16{1, 2, 3}, uint16(6)))
		assert.Equal(t, uint16(6), Product([]uint16{1, 2, 3}))
	})

	t.Run("uint32", func(t *testing.T) {
		assert.True(t, SumsTo([]uint32{1, 2, 3}, uint32(6)))
		assert.Equal(t, uint32(6), Product([]uint32{1, 2, 3}))
	})

	t.Run("uint64", func(t *testing.T) {
		assert.True(t, SumsTo([]uint64{1, 2, 3}, uint64(6)))
		assert.Equal(t, uint64(6), Product([]uint64{1, 2, 3}))
	})

	t.Run("float32", func(t *testing.T) {
		assert.True(t, SumsTo([]float32{1.0, 2.0, 3.0}, float32(6.0)))
		assert.Equal(t, float32(6.0), Product([]float32{1.0, 2.0, 3.0}))
	})
}

