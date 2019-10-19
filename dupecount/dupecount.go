package dupecount

import "strings"

// Dupecount returns a case-incensitive count of alpha-numeric characters
func Dupecount(input string) int {
	input = strings.ToLower(input)
	valmap := make(map[rune]int)
	dupecount := 0

	for _, c := range input {
		if cnt, ok := valmap[c]; !ok {
			valmap[c] = 1
		} else {
			if cnt == 1 {
				dupecount++
			}
			valmap[c]++
		}
	}
	return dupecount
}
