package possum

import "testing"

func TestPositiveSum(t *testing.T) {
	tests := []struct {
		numbers []int
		want    int
	}{
		{
			numbers: []int{-1, 2, 3},
			want:    5,
		},
		{
			numbers: []int{-1, 2, -3},
			want:    2,
		},
		{
			numbers: []int{-1, -2, -3},
			want:    0,
		},
		{
			numbers: []int{1, 2, 3},
			want:    6,
		},
		{
			numbers: []int{},
			want:    0,
		},
	}

	for _, test := range tests {
		if got := PositiveSum(test.numbers); got != test.want {
			t.Errorf("PositiveSum(%v) = %d", test.numbers, got)
		}
	}
}
