package testutils

import (
	"fmt"
	"testing"

	"aoc20-go/utils"
)

func MustReadIntegers(t *testing.T, filename string) []int {
	t.Helper()
	result, err := utils.ReadIntegers(filename)
	if err != nil {
		t.Fatalf("utils.ReadIntegers(%q) error = %v", filename, err)
	}
	return result
}

func MustReadSample(t *testing.T, day string) []int {
	t.Helper()
	filename := fmt.Sprintf("../input/%s-sample.in", day)
	return MustReadIntegers(t, filename)
}

func MustReadReal(t *testing.T, day string) []int {
	t.Helper()
	filename := fmt.Sprintf("../input/%s.in", day)
	return MustReadIntegers(t, filename)
}
