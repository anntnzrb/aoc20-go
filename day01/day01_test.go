package day01

import (
	"testing"

	aocmath "aoc20-go/math"
	"aoc20-go/testutils"
)

func TestSumsTo(t *testing.T) {
	numbers := []int{1721, 299}
	result := aocmath.SumsTo(numbers, 2020)
	testutils.AssertBool(t, result, true, "SumsTo([1721, 299], 2020)")
}

func TestNSum2020(t *testing.T) {
	entries := []int{1721, 979, 366, 299, 675, 1456}
	expected := []int{1721, 299}

	result := NSum(entries, 2020, 2)

	if len(result) != 2 {
		t.Errorf("NSum() length = %d, want 2", len(result))
	}

	if (result[0] != expected[0] || result[1] != expected[1]) && (result[0] != expected[1] || result[1] != expected[0]) {
		t.Errorf("NSum() = %v, want %v", result, expected)
	}
}

func TestProduct(t *testing.T) {
	numbers := []int{1721, 299}
	result := aocmath.Product(numbers)
	testutils.AssertEqual(t, result, 514579, "Product([1721, 299])")
}

func TestThreeSum2020(t *testing.T) {
	entries := []int{1721, 979, 366, 299, 675, 1456}
	result := NSum(entries, 2020, 3)
	product := aocmath.Product(result)
	testutils.AssertEqual(t, product, 241861950, "Product of three numbers that sum to 2020")
}

func TestParseInput(t *testing.T) {
	result := testutils.MustReadSample(t, "day01")
	expected := []int{1721, 979, 366, 299, 675, 1456}
	testutils.AssertSliceEqual(t, result, expected, "Sample input parsing")
}

func TestRealInput(t *testing.T) {
	entries := testutils.MustReadReal(t, "day01")
	numbers := NSum(entries, 2020, 2)
	result := aocmath.Product(numbers)
	testutils.AssertEqual(t, result, 866436, "Product of two numbers that sum to 2020")
}

func TestRealInputThreeSum(t *testing.T) {
	entries := testutils.MustReadReal(t, "day01")
	numbers := NSum(entries, 2020, 3)
	result := aocmath.Product(numbers)
	testutils.AssertEqual(t, result, 276650720, "Product of three numbers that sum to 2020")
}
