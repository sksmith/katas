package splitstrings

import (
	"testing"
)

func TestSplitstring(t *testing.T) {
	tests := []struct {
		value string
		want  []string
	}{
		{
			value: "abc",
			want:  []string{"ab", "c_"},
		},
		{
			value: "abcdef",
			want:  []string{"ab", "cd", "ef"},
		},
	}

	for _, test := range tests {
		got := Splitstrings(test.value)
		if len(got) != len(test.want) {
			t.Errorf("Splitstrings(%s) len=[%d]", test.value, len(got))
			continue
		}
		for i, v := range got {
			if got[i] != test.want[i] {
				t.Errorf("Splitstrings(%s) got[%d]=[%s] want[%d]=[%s]", test.value, i, v, i, test.want[i])
			}
		}
	}
}
