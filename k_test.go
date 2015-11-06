package psl

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
		numbers []byte
		w       []bool
	}{
		{Kombinacia{1, 2, 3, 4, 5}, []byte{0, 1, 2, 3, 4, 5, 6}, []bool{false, true, true, true, true, true, false}},
		{Kombinacia{}, []byte{0, 1, 2}, []bool{false, false, false}},
	}
	for _, test := range tests {
		for i, n := range test.numbers {
			ok := test.k.Contains(n)
			if ok != test.w[i] {
				t.Errorf("Excepted: (%b), Got: (%b)", ok, test.w[i])
			}
		}
	}
}
