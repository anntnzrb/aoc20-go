package math

// Numeric is a constraint that permits any numeric type that supports
// arithmetic operations (+, -, *, /). This includes all signed and unsigned
// integer types as well as floating-point types.
type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
	~float32 | ~float64
}

// SumsTo reports whether the sum of all numbers in the slice equals the target value.
// It works with any numeric type that supports addition.
//
// Example:
//
//	SumsTo([]int{1, 2, 3}, 6)     // returns true
//	SumsTo([]float64{1.5, 2.5}, 4.0) // returns true
//	SumsTo([]int{1, 2, 3}, 7)     // returns false
func SumsTo[T Numeric](numbers []T, target T) bool {
	var sum T
	for _, num := range numbers {
		sum += num
	}
	return sum == target
}

// Product returns the product of all numbers in the slice.
// It works with any numeric type that supports multiplication.
// Returns 1 for an empty slice (multiplicative identity).
//
// Example:
//
//	Product([]int{2, 3, 4})        // returns 24
//	Product([]float64{2.5, 4.0})   // returns 10.0
//	Product([]int{})               // returns 1
func Product[T Numeric](numbers []T) T {
	var result T = 1
	for _, num := range numbers {
		result *= num
	}
	return result
}
