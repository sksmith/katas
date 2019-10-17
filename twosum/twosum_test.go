package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		exp1    int
		exp2    int
	}{
		{
			numbers: []int{1, 2},
			target:  3,
			exp1:    0,
			exp2:    1,
		},
		{
			numbers: []int{1, 2, 2},
			target:  4,
			exp1:    1,
			exp2:    2,
		},
		{
			numbers: []int{2, 1, 2},
			target:  4,
			exp1:    0,
			exp2:    2,
		},
		{
			numbers: []int{2, 2, 2},
			target:  4,
			exp1:    0,
			exp2:    1,
		},
		{
			numbers: []int{2, -1, -1},
			target:  -2,
			exp1:    1,
			exp2:    2,
		},
	}
	for _, test := range tests {
		got1, got2 := TwoSum(test.numbers, test.target)

		if got1 != test.exp1 || got2 != test.exp2 {
			t.Errorf("TwoSum(%v, %d) = %d, %d", test.numbers, test.target, got1, got2)
		}
	}
}
