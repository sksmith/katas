package removetwo

// RemoveNb returns all possible combinations
// of two numbers (a, b) that could be remove from
// 1..n where the product of a and b would equal
// the sum of the rest of the numbers
// The returned array will be sorted by the first
// integer in the pair (a)
func RemoveNb(m uint64) [][2]uint64 {
	var s, i, j, t uint64
	var values [][2]uint64

	for i = 1; i <= m; i++ {
		s += i
	}

	for i = 1; i < m; i++ {
		for j = i + 1; j < m; j++ {
			t = s - i - j
			if t == i*j {
				values = append(values, [2]uint64{i, j})
				values = append(values, [2]uint64{j, i})
			}
		}
	}

	return values
}
