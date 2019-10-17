package snail

import "testing"

func TestSnail(t *testing.T) {
	tests := []struct {
		area [][]int
		want []int
	}{
		{
			area: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
	}
	for _, test := range tests {
		got := Snail(test.area)

		if len(got) == len(test.want) {
			for i, exp := range test.want {
				if got[i] != exp {
					t.Errorf("Snail(%v) = %v", test.area, got)
					break
				}
			}
		} else {
			t.Errorf("Snail(%v) = %v", test.area, got)
		}
	}
}
