package filter

import (
	"testing"

	"github.com/melias122/psl/komb"
)

func TestNtica(t *testing.T) {
	tests := []struct {
		k komb.Kombinacia
		f Filter
		w bool
	}{
		{komb.Kombinacia{1, 2, 3, 4, 5}, NewNtica(5, komb.Tica{0, 0, 0, 0, 1}), true},
		{komb.Kombinacia{1, 2, 3, 4, 6}, NewNtica(5, komb.Tica{1, 0, 0, 1, 0}), true},
		{komb.Kombinacia{1, 2, 3, 5, 6}, NewNtica(5, komb.Tica{0, 1, 1, 0, 0}), true},
		{komb.Kombinacia{1, 2, 3, 5, 7}, NewNtica(5, komb.Tica{2, 0, 1, 0, 0}), true},
		{komb.Kombinacia{1, 2, 4, 5, 7}, NewNtica(5, komb.Tica{1, 2, 0, 0, 0}), true},
		{komb.Kombinacia{1, 2, 4, 6, 8}, NewNtica(5, komb.Tica{3, 1, 0, 0, 0}), true},
		{komb.Kombinacia{1, 3, 5, 7, 9}, NewNtica(5, komb.Tica{5, 0, 0, 0, 0}), true},

		{komb.Kombinacia{1, 2, 3, 4, 5}, NewNtica(5, komb.Tica{5, 0, 0, 0, 0}), false},
	}
	for _, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v)", test.w, ok)
		}
	}
}
