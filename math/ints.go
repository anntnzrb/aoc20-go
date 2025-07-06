package math

func SumsTo(numbers []int, target int) bool {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum == target
}

func Product(numbers []int) int {
	result := 1
	for _, num := range numbers {
		result *= num
	}
	return result
}
