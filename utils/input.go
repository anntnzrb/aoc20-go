package utils

import (
	"os"
	"strconv"
	"strings"
)

func ReadLines(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	return lines, nil
}

func ReadIntegers(filename string) ([]int, error) {
	lines, err := ReadLines(filename)
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
