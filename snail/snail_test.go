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
		{
			area: [][]int{
				{1, 2, 3, 8},
				{4, 5, 6, 3},
				{7, 8, 9, 2},
				{5, 6, 7, 8},
			},
			want: []int{1, 2, 3, 8, 3, 2, 8, 7, 6, 5, 7, 4, 5, 6, 9, 8},
		},
		{
			area: [][]int{
				{1, 2, 3, 4},
				{4, 5, 6, 8},
				{7, 8, 9, 2},
			},
			want: []int{1, 2, 3, 4, 8, 2, 9, 8, 7, 4, 5, 6},
		},
		{
			area: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{4, 4, 4},
			},
			want: []int{1, 2, 3, 6, 9, 4, 4, 4, 7, 4, 5, 8},
		},
		{
			area: [][]int{
				{1},
				{4},
				{7},
				{4},
			},
			want: []int{1, 4, 7, 4},
		},
		{
			area: [][]int{
				{1, 4, 7, 4},
			},
			want: []int{1, 4, 7, 4},
		},
	}
	for testIdx, test := range tests {
		got := Snail(test.area)

		if len(got) == len(test.want) {
			for i, exp := range test.want {
				if got[i] != exp {
					t.Errorf("Test %d: Snail(%v) = %v", testIdx, test.area, got)
					break
				}
			}
		} else {
			t.Errorf("Test %d: Snail(%v) = %v", testIdx, test.area, got)
		}
	}
}
