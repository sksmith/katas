package cubepile

import "testing"

func TestFindNb(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{
			input: 4183059834009,
			want:  2022,
		},
		{
			input: 24723578342962,
			want:  -1,
		},
	}

	for _, test := range tests {
		if got := FindNb(test.input); got != test.want {
			t.Errorf("FindNb(%d) got[%d] want[%d]", test.input, got, test.want)
		}
	}
}
