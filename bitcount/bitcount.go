package bitcount

import "math/bits"

// CountBits returns the number of ones in the bit version of the supplied uint
func CountBits(n uint) int {
	return bits.OnesCount(n)
}
