package day01

import "aoc20-go/utils"

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
	return utils.ReadIntegers(filename)
}
