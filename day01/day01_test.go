package day01

import "testing"

func TestSumsTo(t *testing.T) {
	a, b, target := 1721, 299, 2020
	expected := true

	result := SumsTo(a, b, target)

	if result != expected {
		t.Errorf("SumsTo(%d, %d, %d) = %v, want %v", a, b, target, result, expected)
	}
}

func TestTwoSum2020(t *testing.T) {
	entries := []int{1721, 979, 366, 299, 675, 1456}
	expectedA, expectedB := 1721, 299

	a, b := TwoSum(entries, 2020)

	if (a != expectedA || b != expectedB) && (a != expectedB || b != expectedA) {
		t.Errorf("TwoSum() = %d, %d, want %d, %d", a, b, expectedA, expectedB)
	}
}

func TestParseInput(t *testing.T) {
	expected := []int{1721, 979, 366, 299, 675, 1456}

	result, err := ParseInput("../input/day01-sample.in")

	if err != nil {
		t.Fatalf("ParseInput() error = %v", err)
	}

	if len(result) != len(expected) {
		t.Errorf("ParseInput() length = %d, want %d", len(result), len(expected))
	}

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("ParseInput()[%d] = %d, want %d", i, result[i], v)
		}
	}
}

func TestRealInput(t *testing.T) {
	entries, err := ParseInput("../input/day01.in")
	if err != nil {
		t.Fatalf("ParseInput() error = %v", err)
	}

	a, b := TwoSum(entries, 2020)
	result := a * b
	expected := 866436

	if result != expected {
		t.Errorf("Product of numbers that sum to 2020 = %d, want %d", result, expected)
	}
}
