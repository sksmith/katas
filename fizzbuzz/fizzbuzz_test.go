package fizzbuzz

import "testing"

const errm = "FizzBuzz(%d)\n   got=%v\n  want=%v\n"

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		value int
		want  []string
	}{
		{
			value: 15,
			want: []string{
				"1",
				"2",
				"Fizz",
				"4",
				"Buzz",
				"Fizz",
				"7",
				"8",
				"Fizz",
				"Buzz",
				"11",
				"Fizz",
				"13",
				"14",
				"FizzBuzz"},
		},
	}
	for _, test := range tests {
		got := FizzBuzz(test.value)
		if len(got) != len(test.want) {
			t.Errorf(errm, test.value, got, test.want)
			continue
		}
		for i := 0; i < len(got); i++ {
			if got[i] != test.want[i] {
				t.Errorf(errm, test.value, got, test.want)
				break
			}
		}
	}
}
