package day01

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
