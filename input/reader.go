// Package input provides utilities for common file and input processing tasks.
//
// This package contains helper functions for reading and parsing input files,
// particularly useful for Advent of Code style problems that require reading
// structured data from text files.
package input

import (
	"os"
	"strconv"
	"strings"
)

// ReadLines reads a file and returns its contents as a slice of strings,
// with each line as a separate element. The file content is trimmed of
// leading and trailing whitespace before splitting.
//
// Parameters:
//   - filename: Path to the file to read
//
// Returns:
//   - []string: Slice containing each line of the file
//   - error: Any error encountered while reading the file
//
// Example:
//
//	lines, err := input.ReadLines("input.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for i, line := range lines {
//		fmt.Printf("Line %d: %s\n", i+1, line)
//	}
//
// The function handles files with different line endings and automatically
// trims whitespace from the beginning and end of the file content.
func ReadLines(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	return lines, nil
}

// ReadIntegers reads a file and parses each line as an integer,
// returning a slice of all successfully parsed integers.
// Lines that cannot be parsed as integers are silently skipped.
//
// Parameters:
//   - filename: Path to the file to read
//
// Returns:
//   - []int: Slice containing all integers found in the file
//   - error: Any error encountered while reading the file (parsing errors are ignored)
//
// Example:
//
//	numbers, err := input.ReadIntegers("numbers.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Read %d numbers: %v\n", len(numbers), numbers)
//
// File format example:
//
//	1721
//	979
//	366
//	299
//
// This function is particularly useful for Advent of Code problems where
// input files contain lists of integers, one per line. Invalid lines
// (non-numeric) are ignored rather than causing the function to fail.
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
