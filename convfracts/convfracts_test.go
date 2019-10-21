package convfracts

import "testing"

func TestConvertFracts(t *testing.T) {
	tests := []struct {
		input [][]int
		want  string
	}{
		{
			input: [][]int{{1, 2}, {1, 3}, {1, 4}},
			want:  "(6,12)(4,12)(3,12)",
		},
		{
			input: [][]int{{69, 130}, {87, 1310}, {30, 40}},
			want:  "(18078,34060)(2262,34060)(25545,34060)",
		},
	}

	for i, test := range tests {
		if got := ConvertFracts(test.input); got != test.want {
			t.Errorf("ConvertFracts() test=[%d] want[%s] got[%s]", i, test.want, got)
		}
	}
}

func TestCalcNumerator(t *testing.T) {
	tests := []struct {
		numer int
		denom int
		lcm   int
		want  int
	}{
		{
			numer: 1,
			denom: 12,
			lcm:   12,
			want:  1,
		},
		{
			numer: 2,
			denom: 12,
			lcm:   6,
			want:  1,
		},
		{
			numer: 1,
			denom: 2,
			lcm:   12,
			want:  6,
		},
	}

	for _, test := range tests {
		if got := CalcNumerator(test.numer, test.denom, test.lcm); got != test.want {
			t.Errorf("ConvertFracts(%d, %d, %d) got[%d] want[%d]", test.numer, test.denom, test.lcm, got, test.want)
		}
	}
}

func TestSimplify(t *testing.T) {
	tests := []struct {
		frac []int
		want []int
	}{
		{
			frac: []int{1, 2},
			want: []int{1, 2},
		},
		{
			frac: []int{2, 4},
			want: []int{1, 2},
		},
		{
			frac: []int{4, 10},
			want: []int{2, 5},
		},
		{
			frac: []int{30, 40},
			want: []int{3, 4},
		},
	}

	for _, test := range tests {
		if got := Simplify(test.frac); got[0] != test.want[0] || got[1] != test.want[1] {
			t.Errorf("ConvertFracts(%v) got[%v] want[%v]", test.frac, got, test.want)
		}
	}
}
