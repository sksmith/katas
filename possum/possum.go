package possum

// PositiveSum returns the sum of positive values in the supplied array
func PositiveSum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		if val > 0 {
			sum += val
		}
	}
	return sum
}
