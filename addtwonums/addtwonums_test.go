package addtwonums

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1   []int
		l2   []int
		want []int
	}{
		{
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
		{
			l1:   []int{0},
			l2:   []int{0},
			want: []int{0},
		},
		{
			l1:   []int{1, 8},
			l2:   []int{0},
			want: []int{1, 8},
		},
		{
			l1:   []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			l2:   []int{5, 6, 4},
			want: []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
		{
			l1:   []int{5},
			l2:   []int{5},
			want: []int{0, 1},
		},
		{
			l1:   []int{1},
			l2:   []int{9, 9},
			want: []int{0, 0, 1},
		},
	}
	for _, test := range tests {
		l1 := NewListNode(test.l1)
		l2 := NewListNode(test.l2)

		got := addTwoNumbersNoMul(l1, l2)

		s := NewListNode(test.want)
		e := s
		a := got
		for e != nil {
			if a == nil {
				t.Errorf("addTwoNumbers(%v, %v)\n   got: %v\n  want: %v\n", test.l1, test.l2, got, s)
			}
			if e.Val != a.Val {
				t.Errorf("addTwoNumbers(%v, %v)\n   got: %v\n  want: %v\n", test.l1, test.l2, got, s)
			}
			e = e.Next
			a = a.Next
		}
	}
}
