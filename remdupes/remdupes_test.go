package remdupes

import "testing"

func TestRemDupes(t *testing.T) {
	var tests = []struct {
		a    []int
		want []int
	}{
		{
			a:    []int{0, 1, 2, 3},
			want: []int{0, 1, 2, 3},
		},
		{
			a:    []int{0, 0, 1, 2},
			want: []int{0, 1, 2},
		},
		{
			a:    []int{0, 1, 2, 2},
			want: []int{0, 1, 2},
		},
		{
			a:    []int{1, 2, 1, 1, 3},
			want: []int{1, 2, 3},
		},
		{
			a:    []int{5, 5, 5, 5, 5},
			want: []int{5},
		},
		{
			a:    []int{},
			want: []int{},
		},
		{
			a:    []int{3, 1, 2, 1, 1},
			want: []int{3, 1, 2},
		},
	}

	for i, test := range tests {
		got := RemDupes(test.a)

		if len(got) != len(test.want) {
			t.Errorf("test %d:\n   got: %v\n  want: %v", i, got, test.want)
			continue
		}

		for i, v := range got {
			if test.want[i] != v {
				t.Errorf("test %d:\n   got: %v\n  want: %v", i, got, test.want)
				continue
			}
		}
	}
}

func TestRemDupesSorted(t *testing.T) {
	var tests = []struct {
		a    []int
		want []int
	}{
		{
			a:    []int{3, 0, 1, 2},
			want: []int{0, 1, 2, 3},
		},
		{
			a:    []int{1, 2, 0, 0},
			want: []int{0, 1, 2},
		},
		{
			a:    []int{0, 2, 1, 2},
			want: []int{0, 1, 2},
		},
		{
			a:    []int{1, 2, 1, 1, 3},
			want: []int{1, 2, 3},
		},
		{
			a:    []int{5, 5, 5, 5, 5},
			want: []int{5},
		},
		{
			a:    []int{},
			want: []int{},
		},
	}

	for i, test := range tests {
		got := RemDupesSorted(test.a)

		if len(got) != len(test.want) {
			t.Errorf("test %d:\n   got: %v\n  want: %v", i, got, test.want)
			continue
		}

		for i, v := range got {
			if test.want[i] != v {
				t.Errorf("test %d:\n   got: %v\n  want: %v", i, got, test.want)
				continue
			}
		}
	}
}
