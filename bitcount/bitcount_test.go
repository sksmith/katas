package bitcount

import "testing"

func TestCountBits(t *testing.T) {
	var tests = []struct {
		input uint
		want  int
	}{
		{0, 0},
		{4, 1},
		{7, 3},
		{9, 2},
		{10, 2},
	}

	for _, test := range tests {
		if got := CountBits(test.input); got != test.want {
			t.Errorf("CountBits(%d) = %v", test.input, got)
		}
	}
}
