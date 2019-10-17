package twosum

// TwoSum retrieves the first two indexes that sum to the target
func TwoSum(numbers []int, target int) (idx1, idx2 int) {
	val1 := 0
	for idx1, val1 = range numbers {
		for idx2 = idx1 + 1; idx2 < len(numbers); idx2++ {
			if val1+numbers[idx2] == target {
				return
			}
		}
	}
	return -1, -1
}
