package fizzbuzz

import "strconv"

// FizzBuzz prints
// Fizz for multiples of three
// Buzz for multiples of five
// FizzBuzz for both
func FizzBuzz(c int) []string {
	result := make([]string, c)

	for i := 1; i <= c; i++ {
		var val string

		if i%3 != 0 && i%5 != 0 {
			val = strconv.Itoa(i)
		} else {
			if i%3 == 0 {
				val = "Fizz"
			}
			if i%5 == 0 {
				val += "Buzz"
			}
		}
		result[i-1] = val
	}

	return result
}
