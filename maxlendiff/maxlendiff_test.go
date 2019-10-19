package maxlendiff

import "testing"

func TestMaxlendiff(t *testing.T) {
	tests := []struct {
		a1   []string
		a2   []string
		want int
	}{
		{
			a1:   []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"},
			a2:   []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"},
			want: 13,
		},
		{
			a1:   []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"},
			a2:   []string{"bbbaaayddqbbrrrv"},
			want: 10,
		},
		{
			a1:   []string{},
			a2:   []string{"bbbaaayddqbbrrrv"},
			want: -1,
		},
		{
			a1:   []string{"asdfasdf"},
			a2:   []string{},
			want: -1,
		},
	}

	for i, test := range tests {
		if got := Maxlendiff(test.a1, test.a2); got != test.want {
			t.Errorf("Maxlendiff() test[%d] got[%d] want[%d]", i, got, test.want)
		}
	}
}
