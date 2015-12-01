package psl

import (
	"testing"
)

func TestFilterR(t *testing.T) {
	cisla := Nums{
		NewNum(1, 5, 35),
		NewNum(2, 5, 35),
		NewNum(3, 5, 35),
	}
	tests := []struct {
		k Kombinacia
		f Filter
		w bool
	}{
		{Kombinacia{1, 2, 3}, NewFilterR1(0.0, 1.0, cisla, 3), true},
		{Kombinacia{1, 2}, NewFilterR1(0, -0.99, cisla, 3), false},
		{Kombinacia{1, 2}, NewFilterR1(0.1, 1, cisla, 3), true},
		{Kombinacia{1, 2, 3}, NewFilterR1(0.101, 1, cisla, 3), false},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v) (%v)", test.w, ok, test)
		}
	}
}
