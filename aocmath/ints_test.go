package aocmath

import (
	"testing"

	"aoc20-go/testutil"
	"github.com/stretchr/testify/assert"
)

func TestSumsTo(t *testing.T) {
	type input struct {
		numbers []int
		target  int
	}

	testutil.Run(t,
		func(in input) bool { return SumsTo(in.numbers, in.target) },
		[]struct {
			Name     string
			Input    input
			Expected bool
		}{
			testutil.T("basic sum matches", input{[]int{1, 2, 3}, 6}, true),
			testutil.T("basic sum doesn't match", input{[]int{1, 2, 3}, 7}, false),
			testutil.T("empty slice", input{[]int{}, 0}, true),
			testutil.T("empty slice non-zero target", input{[]int{}, 5}, false),
			testutil.T("single element matches", input{[]int{42}, 42}, true),
			testutil.T("single element doesn't match", input{[]int{42}, 43}, false),
			testutil.T("negative numbers", input{[]int{-1, -2, -3}, -6}, true),
			testutil.T("mixed positive and negative", input{[]int{10, -5, 2, -3}, 4}, true),
			testutil.T("zero sum", input{[]int{5, -5, 3, -3}, 0}, true),
			testutil.T("large numbers", input{[]int{1000000, 2000000, 3000000}, 6000000}, true),
		})
}

func TestSumsToFloat64(t *testing.T) {
	type input struct {
		numbers []float64
		target  float64
	}

	testutil.Run(t,
		func(in input) bool { return SumsTo(in.numbers, in.target) },
		[]struct {
			Name     string
			Input    input
			Expected bool
		}{
			testutil.T("basic float sum", input{[]float64{1.5, 2.5, 3.0}, 7.0}, true),
			testutil.T("decimal precision", input{[]float64{0.25, 0.25, 0.5}, 1.0}, true),
			testutil.T("negative floats", input{[]float64{-1.5, -2.5}, -4.0}, true),
			testutil.T("mixed float values", input{[]float64{1.1, -0.1, 2.0}, 3.0}, true),
		})
}

func TestSumsToUint(t *testing.T) {
	type input struct {
		numbers []uint
		target  uint
	}

	testutil.Run(t,
		func(in input) bool { return SumsTo(in.numbers, in.target) },
		[]struct {
			Name     string
			Input    input
			Expected bool
		}{
			testutil.T("basic uint sum", input{[]uint{1, 2, 3}, uint(6)}, true),
			testutil.T("large uint values", input{[]uint{4294967295, 1}, uint(4294967296)}, true),
		})
}

func TestProduct(t *testing.T) {
	testutil.Run(t,
		func(numbers []int) int { return Product(numbers) },
		[]struct {
			Name     string
			Input    []int
			Expected int
		}{
			testutil.T("basic product", []int{2, 3, 4}, 24),
			testutil.T("empty slice", []int{}, 1),
			testutil.T("single element", []int{42}, 42),
			testutil.T("with zero", []int{1, 2, 0, 4}, 0),
			testutil.T("negative numbers even count", []int{-2, -3}, 6),
			testutil.T("negative numbers odd count", []int{-2, -3, -1}, -6),
			testutil.T("mixed positive and negative", []int{2, -3, 4}, -24),
			testutil.T("ones", []int{1, 1, 1, 1}, 1),
			testutil.T("large numbers", []int{100, 200}, 20000),
		})
}

func TestProductFloat64(t *testing.T) {
	testutil.Run(t,
		func(numbers []float64) float64 { return Product(numbers) },
		[]struct {
			Name     string
			Input    []float64
			Expected float64
		}{
			testutil.T("basic float product", []float64{2.5, 4.0}, 10.0),
			testutil.T("decimal values", []float64{0.5, 0.2, 10.0}, 1.0),
			testutil.T("negative floats", []float64{-2.0, 3.0}, -6.0),
			testutil.T("with zero", []float64{1.5, 0.0, 2.0}, 0.0),
		})
}

func TestProductUint(t *testing.T) {
	testutil.Run(t,
		func(numbers []uint) uint { return Product(numbers) },
		[]struct {
			Name     string
			Input    []uint
			Expected uint
		}{
			testutil.T("basic uint product", []uint{2, 3, 4}, uint(24)),
			testutil.T("large uint values", []uint{1000, 1000}, uint(1000000)),
		})
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
