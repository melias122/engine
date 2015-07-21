package filter

import (
	"testing"

	"github.com/melias122/psl/komb"
)

func TestZhoda(t *testing.T) {
	tests := []struct {
		k komb.Kombinacia
		f Filter
		w bool
	}{
		{komb.Kombinacia{1}, NewZhoda(3, 2, 3, komb.Kombinacia{1}), true},
		{komb.Kombinacia{1}, NewZhoda(3, 0, 0, komb.Kombinacia{1}), false},
		{komb.Kombinacia{1}, NewZhoda(3, 0, 3, komb.Kombinacia{1}), true},
		{komb.Kombinacia{1}, NewZhoda(3, 0, 3, komb.Kombinacia{2}), true},
		{komb.Kombinacia{1, 2, 3}, NewZhoda(3, 2, 2, komb.Kombinacia{1, 2}), true},
		{komb.Kombinacia{1, 2, 3}, NewZhoda(3, 0, 2, komb.Kombinacia{1, 2, 3}), false},
		{komb.Kombinacia{1, 2, 3}, NewZhoda(3, 0, 3, komb.Kombinacia{4, 5, 6}), true},
		{komb.Kombinacia{1, 2, 3}, NewZhoda(3, 1, 3, komb.Kombinacia{4, 5, 6}), false},

		{komb.Kombinacia{1, 2, 3}, NewZhoda(3, 0, 3, komb.Kombinacia{1, 2, 3}), true},
		{komb.Kombinacia{1, 2, 3}, NewZhoda(3, 1, 3, komb.Kombinacia{1, 2, 3}), true},
		{komb.Kombinacia{1, 2, 3}, NewZhoda(3, 2, 3, komb.Kombinacia{1, 2, 3}), true},
		{komb.Kombinacia{1, 2, 3}, NewZhoda(3, 3, 3, komb.Kombinacia{1, 2, 3}), true},

		{komb.Kombinacia{1, 2, 3}, NewZhoda(3, 0, 2, komb.Kombinacia{1, 2, 3}), false},
		{komb.Kombinacia{1, 2, 3}, NewZhoda(3, 0, 1, komb.Kombinacia{1, 2, 3}), false},
		{komb.Kombinacia{1, 2, 3}, NewZhoda(3, 0, 0, komb.Kombinacia{1, 2, 3}), false},
	}
	for i, test := range tests {
		ok := test.f.Check(test.k)
		if ok != test.w {
			t.Errorf("Excepted: (%v), Got: (%v) (test %d)", test.w, ok, i+1)
		}
	}
}
