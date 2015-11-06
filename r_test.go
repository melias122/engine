package psl

import (
	"testing"

	// "github.com/melias122/psl/komb"
	// "github.com/melias122/psl/num"
)

func TestR(t *testing.T) {
	cisla := num.Nums{
		num.New(1, 5, 35),
		num.New(2, 5, 35),
		num.New(3, 5, 35),
	}
	tests := []struct {
		k komb.Kombinacia
		f Filter
		w bool
	}{
		{komb.Kombinacia{1, 2, 3}, NewR(3, 0.0, 1.0, cisla, ""), true},
		{komb.Kombinacia{1, 2}, NewR(3, 0, -0.99, cisla, ""), false},
		{komb.Kombinacia{1, 2}, NewR(3, 0.1, 1, cisla, ""), true},
		{komb.Kombinacia{1, 2, 3}, NewR(3, 0.101, 1, cisla, ""), false},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v) (%v)", test.w, ok, test)
		}
	}
}
