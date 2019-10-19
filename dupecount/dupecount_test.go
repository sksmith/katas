package dupecount

import (
	"testing"
)

func TestDupecount(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "abcde",
			want:  0,
		},
		{
			input: "aabbcde",
			want:  2,
		},
		{
			input: "aabBcde",
			want:  2,
		},
		{
			input: "indivisibility",
			want:  1,
		},
		{
			input: "Indivisibilities",
			want:  2,
		},
		{
			input: "aA11",
			want:  2,
		},
		{
			input: "ABBA",
			want:  2,
		},
	}

	for _, test := range tests {
		if got := Dupecount(test.input); got != test.want {
			t.Errorf("Dupecount(%s) = got[%d] want[%d]", test.input, got, test.want)
		}
	}
}
