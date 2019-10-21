package splitstrings

import "math"

// Splitstrings splits a string into pairs of two characters,
// if there are an odd number of characters adds a trailing '_'
// https://www.codewars.com/kata/515de9ae9dcfc28eb6000001/train/go
func Splitstrings(str string) (split []string) {
	flen := float64(len(str)) / 2.0
	round := math.Round(flen)
	split = make([]string, int(round))
	for i := 0; i < len(split); i++ {
		x := i * 2
		if x+1 < len(str) {
			split[i] = str[x : x+2]
		} else {
			split[i] = str[x:x+1] + "_"
		}
	}
	return split
}
