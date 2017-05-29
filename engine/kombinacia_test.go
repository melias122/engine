package engine

import "testing"

func TestString(t *testing.T) {
	k := Kombinacia{1, 2, 3, 4, 5}
	if k.String() != "1 2 3 4 5" {
		t.Errorf("Excepted: (1 2 3 4 5), Got: (%s)", k.String())
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		k       Kombinacia
		numbers []int
		w       []bool
	}{
		{Kombinacia{1, 2, 3, 4, 5}, []int{0, 1, 2, 3, 4, 5, 6}, []bool{false, true, true, true, true, true, false}},
		{Kombinacia{}, []int{0, 1, 2}, []bool{false, false, false}},
	}
	for _, test := range tests {
		for i, n := range test.numbers {
			ok := test.k.Contains(n)
			if ok != test.w[i] {
				t.Errorf("Excepted: (%v), Got: (%v)", ok, test.w[i])
			}
		}
	}
}