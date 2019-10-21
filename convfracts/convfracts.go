package convfracts

import "fmt"

func ConvertFracts(a [][]int) string {
	denoms := make([]int, len(a))

	for i := range a {
		a[i] = Simplify(a[i])
		denoms[i] = a[i][1]
	}

	lcm := LCM(denoms[0], denoms[1], denoms[2:]...)

	val := ""
	for _, v := range a {
		numerator := CalcNumerator(v[0], v[1], lcm)
		// fmt.Printf("CalcNumerator(%d, %d, %d) == %d\n", v[0], v[1], lcm, numerator)
		val = fmt.Sprintf(val+"(%d,%d)", numerator, lcm)
	}

	return val
}

// Simplify returns a simplified fraction
func Simplify(frac []int) []int {
	gcd := GCD(frac[0], frac[1])
	return []int{frac[0] / gcd, frac[1] / gcd}
}

// CalcNumerator determines what the numerator should be based off of the denominator and its new value (the lcm)
func CalcNumerator(numerator, denominator, lcm int) int {
	multiple := float64(lcm) / float64(denominator)

	return int(float64(numerator) * multiple)
}

// GCD returns greatest common divisor (GCD)
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM returns Least Common Multiple (LCM)
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
