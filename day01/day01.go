package day01

import (
	"aoc20-go/utils"
	"strconv"
)

func SumsTo(a, b, target int) bool {
	return a+b == target
}

func TwoSum(entries []int, target int) (int, int) {
	seen := make(map[int]bool)
	for _, entry := range entries {
		complement := target - entry
		if seen[complement] {
			return entry, complement
		}
		seen[entry] = true
	}
	return 0, 0
}

func ParseInput(filename string) ([]int, error) {
	lines, err := utils.ReadLines(filename)
	if err != nil {
		return nil, err
	}

	numbers := make([]int, 0, len(lines))
	for _, line := range lines {
		if num, err := strconv.Atoi(line); err == nil {
			numbers = append(numbers, num)
		}
	}

	return numbers, nil
}
