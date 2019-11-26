package removetwo

import (
	"log"
	"testing"
)

func BenchmarkRemoveNb(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveNb(4000)
	}
}

func TestFindNbs(t *testing.T) {
	for i := uint64(2); i < 2000; i++ {
		v := RemoveNb(i)
		if v != nil {
			log.Printf("%d: %v", i, v)
		}
	}
}

func TestRemoveNb(t *testing.T) {
	tests := []struct {
		m    uint64
		want [][2]uint64
	}{
		{m: 26, want: [][2]uint64{{15, 21}, {21, 15}}},
		{m: 100, want: nil},
		{m: 101, want: [][2]uint64{{55, 91}, {91, 55}}},
		{m: 102, want: [][2]uint64{{70, 73}, {73, 70}}},
		{m: 110, want: [][2]uint64{{70, 85}, {85, 70}}},
		{m: 10, want: [][2]uint64{{6, 7}, {7, 6}}},
		{m: 1093, want: [][2]uint64{{631, 945}, {945, 631}, {687, 868}, {868, 687}}},
	}

	for i, test := range tests {
		got := RemoveNb(test.m)

		if len(got) != len(test.want) {
			t.Errorf("%2d-RemoveNb(%d) got[%v] want[%v]", i, test.m, got, test.want)
			continue
		}

		for i := range got {
			if got[i][0] != test.want[i][0] || got[i][1] != test.want[i][1] {
				t.Errorf("%2d-RemoveNb(%d) got[%v] want[%v]", i, test.m, got, test.want)
				break
			}
		}
	}
}
