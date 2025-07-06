package day01

import "testing"

func TestSumsTo(t *testing.T) {
	numbers := []int{1721, 299}
	target := 2020
	expected := true

	result := SumsTo(numbers, target)

	if result != expected {
		t.Errorf("SumsTo(%v, %d) = %v, want %v", numbers, target, result, expected)
	}
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
	expected := 514579

	result := Product(numbers)

	if result != expected {
		t.Errorf("Product(%v) = %d, want %d", numbers, result, expected)
	}
}

func TestThreeSum2020(t *testing.T) {
	entries := []int{1721, 979, 366, 299, 675, 1456}
	expected := 241861950

	result := NSum(entries, 2020, 3)
	product := Product(result)

	if product != expected {
		t.Errorf("Product of three numbers that sum to 2020 = %d, want %d", product, expected)
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

	numbers := NSum(entries, 2020, 2)
	result := Product(numbers)
	expected := 866436

	if result != expected {
		t.Errorf("Product of numbers that sum to 2020 = %d, want %d", result, expected)
	}
}

func TestRealInputThreeSum(t *testing.T) {
	entries, err := ParseInput("../input/day01.in")
	if err != nil {
		t.Fatalf("ParseInput() error = %v", err)
	}

	numbers := NSum(entries, 2020, 3)
	result := Product(numbers)
	expected := 276650720

	if result != expected {
		t.Errorf("Product of three numbers that sum to 2020 = %d, want %d", result, expected)
	}
}
