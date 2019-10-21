package convfracts

import "fmt"

func ConvertFracts(a [][]int) string {
	lcm := LCM(1, a[0], a[1:]...)

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

// LCM simplifies the incoming fractions and
// returns Least Common Multiple (LCM)
func LCM(lcm int, a []int, integers ...[]int) int {
	simple := Simplify(a)
	a[0] = simple[0]
	a[1] = simple[1]

	result := lcm * a[1] / GCD(lcm, a[1])

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
