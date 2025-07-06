package day01

import (
	"testing"

	aocmath "aoc20-go/math"
	"aoc20-go/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSumsTo(t *testing.T) {
	numbers := []int{1721, 299}
	result := aocmath.SumsTo(numbers, 2020)
	assert.True(t, result, "SumsTo([1721, 299], 2020)")
}

func TestNSum2020(t *testing.T) {
	entries := []int{1721, 979, 366, 299, 675, 1456}
	expected := []int{1721, 299}

	result := NSum(entries, 2020, 2)

	assert.Len(t, result, 2, "NSum() should return 2 elements")
	assert.ElementsMatch(t, expected, result, "NSum() should return the expected values")
}

func TestProduct(t *testing.T) {
	numbers := []int{1721, 299}
	result := aocmath.Product(numbers)
	assert.Equal(t, 514579, result, "Product([1721, 299])")
}

func TestThreeSum2020(t *testing.T) {
	entries := []int{1721, 979, 366, 299, 675, 1456}
	result := NSum(entries, 2020, 3)
	product := aocmath.Product(result)
	assert.Equal(t, 241861950, product, "Product of three numbers that sum to 2020")
}

func TestParseInput(t *testing.T) {
	filename := "../input/day01-sample.in"
	result, err := utils.ReadIntegers(filename)
	require.NoError(t, err, "Failed to read sample input")
	expected := []int{1721, 979, 366, 299, 675, 1456}
	assert.Equal(t, expected, result, "Sample input parsing")
}

func TestRealInput(t *testing.T) {
	filename := "../input/day01.in"
	entries, err := utils.ReadIntegers(filename)
	require.NoError(t, err, "Failed to read real input")
	numbers := NSum(entries, 2020, 2)
	result := aocmath.Product(numbers)
	assert.Equal(t, 866436, result, "Product of two numbers that sum to 2020")
}

func TestRealInputThreeSum(t *testing.T) {
	filename := "../input/day01.in"
	entries, err := utils.ReadIntegers(filename)
	require.NoError(t, err, "Failed to read real input")
	numbers := NSum(entries, 2020, 3)
	result := aocmath.Product(numbers)
	assert.Equal(t, 276650720, result, "Product of three numbers that sum to 2020")
}
