package remdupes

import "sort"

// RemDupes removes duplicates from an array
func RemDupes(a []int) (b []int) {
	size := len(a)

	for i := 0; i < len(a); i++ {
		if a[i] == -1 {
			continue
		}

		for n := i + 1; n < len(a); n++ {
			if a[n] == -1 {
				continue
			}
			if a[i] == a[n] {
				a[n] = -1
				size--
			}
		}
	}

	result := make([]int, size)
	resIdx := 0

	for i := 0; i < len(a); i++ {
		if a[i] != -1 {
			result[resIdx] = a[i]
			resIdx++
		}
	}
	return result
}

func RemDupesSorted(a []int) (result []int) {
	if len(a) == 0 {
		return a
	}

	sort.Ints(a)

	prev := a[0]
	size := len(a)

	for i := 1; i < len(a); i++ {
		if a[i] == prev {
			a[i] = -1
			size--
		} else {
			prev = a[i]
		}
	}

	result = make([]int, size)
	resIdx := 0
	for i := 0; i < len(a); i++ {
		if a[i] != -1 {
			result[resIdx] = a[i]
			resIdx++
		}
	}
	return result
}
