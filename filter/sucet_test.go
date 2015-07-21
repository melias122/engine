package filter

import (
	"testing"

	"github.com/melias122/psl/komb"
)

func TestSucet(t *testing.T) {
	tests := []struct {
		k komb.Kombinacia
		f Filter
		w bool
	}{
		{komb.Kombinacia{1}, NewSucet(3, 14, 14), true},
		{komb.Kombinacia{1, 11}, NewSucet(3, 14, 14), true},
		{komb.Kombinacia{1, 2, 10}, NewSucet(3, 14, 14), false},
		{komb.Kombinacia{1, 2, 11}, NewSucet(3, 14, 14), true},
		{komb.Kombinacia{1, 2, 3, 4, 5}, NewSucet(5, 0, 14), false},
		{komb.Kombinacia{1, 2, 3, 4, 5}, NewSucet(5, 0, 15), true},
		{komb.Kombinacia{1, 2, 3, 4, 5}, NewSucet(5, 15, 15), true},
		{komb.Kombinacia{1, 2, 3, 4, 5}, NewSucet(5, 15, 55), true},
		{komb.Kombinacia{1, 2, 3, 4, 5}, NewSucet(5, 16, 16), false},

		{komb.Kombinacia{1, 2, 3, 4, 5}, NewSucet(5, 30, 50), false},
		{komb.Kombinacia{1, 2, 3, 4, 25}, NewSucet(5, 30, 50), true},
		{komb.Kombinacia{1, 2, 3, 4, 45}, NewSucet(5, 30, 50), false},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}
