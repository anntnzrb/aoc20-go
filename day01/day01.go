package day01

import "aoc20-go/utils"

func SumsTo(numbers []int, target int) bool {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum == target
}

func NSum(entries []int, target int, n int) []int {
	if n == 1 {
		for _, entry := range entries {
			if entry == target {
				return []int{entry}
			}
		}
		return nil
	}

	for i, num := range entries {
		if result := NSum(entries[i+1:], target-num, n-1); result != nil {
			return append([]int{num}, result...)
		}
	}
	return nil
}

func Product(numbers []int) int {
	result := 1
	for _, num := range numbers {
		result *= num
	}
	return result
}

func ParseInput(filename string) ([]int, error) {
	return utils.ReadIntegers(filename)
}
