package day01

// Entries represents a list of expense report entries to analyze
type Entries []int

// Target represents the target sum we're looking for (e.g., 2020)
type Target int

// Count represents how many numbers should sum to the target
type Count int

// Combination represents a mathematical combination of numbers that sum to the target
type Combination []int

// NSum finds n numbers in entries that sum to target
func NSum(entries Entries, target Target, n Count) Combination {
	if n == 1 {
		for _, entry := range entries {
			if Target(entry) == target {
				return Combination{entry}
			}
		}
		return nil
	}

	for i, num := range entries {
		if result := NSum(entries[i+1:], target-Target(num), n-1); result != nil {
			return append(Combination{num}, result...)
		}
	}
	return nil
}
