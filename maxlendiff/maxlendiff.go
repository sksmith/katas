package maxlendiff

// Maxlendiff returns the maximum difference between any two strings in two arrays
func Maxlendiff(a, b []string) int {
	if a == nil || b == nil || len(a) == 0 || len(b) == 0 {
		return -1
	}

	maxa, mina := maxmin(a)
	maxb, minb := maxmin(b)

	diffa := maxa - minb
	diffb := maxb - mina

	if diffa > diffb {
		return diffa
	}

	return diffb
}

func maxmin(a []string) (max, min int) {
	max = len(a[0])
	min = max

	for _, s := range a {
		l := len(s)
		if max < l {
			max = l
		}
		if l < min {
			min = l
		}
	}

	return
}
